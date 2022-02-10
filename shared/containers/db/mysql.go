package db

import (
	"containers"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gorm_logger "logger/gorm"
)

type MySQLContainer struct {
	opts   containers.Opts
	mysqlC testcontainers.Container
	ctx    *context.Context
}

func NewMysqlContainer() *MySQLContainer {
	mysqlOpts := containers.Opts{
		User:     "forum",
		Password: "secret",
		Database: "gforum",
	}
	return &MySQLContainer{
		opts: mysqlOpts,
	}
}

func (m *MySQLContainer) CreateContainer() {
	log.Info().Msg("setup MySQL Container")
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "mysql:latest",
		ExposedPorts: []string{"3306/tcp"},
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": m.opts.Password,
			"MYSQL_USER":          m.opts.User,
			"MYSQL_PASSWORD":      m.opts.Password,
			"MYSQL_DATABASE":      m.opts.Database,
		},
		WaitingFor: wait.ForLog("port: 3306  MySQL Community Server - GPL"),
	}

	mysqlC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	if err != nil {
		log.Fatal().Msgf("error starting mysql container: %s", err)
	}
	m.mysqlC = mysqlC
	m.ctx = &ctx
}

func (m *MySQLContainer) DbManager() *gorm.DB {
	var db *gorm.DB
	dsn := m.BuildDSN()

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gorm_logger.New(),
	})
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to Open database: %s", err)
	}

	return db
}

func (m *MySQLContainer) BuildDSN() string {
	host, _ := m.mysqlC.Host(*m.ctx)
	p, _ := m.mysqlC.MappedPort(*m.ctx, "3306/tcp")
	port := p.Int()
	defaultDsn := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true"
	dsn := fmt.Sprintf(defaultDsn, m.opts.User, m.opts.Password, host, port, m.opts.Database)
	log.Info().Msgf("Connecting to mysql: %s", dsn)
	return dsn
}

func (m *MySQLContainer) SqlDbManager() *sql.DB {
	var db *sql.DB
	dsn := m.BuildDSN()
	boil.DebugMode = true
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to Open database: %s", err)
	}
	return db
}

func (m *MySQLContainer) CloseContainer() {
	log.Info().Msg("terminating container")
	err := m.mysqlC.Terminate(*m.ctx)
	if err != nil {
		log.Fatal().Msgf("error terminating mysql container: %s", err)
	}
}
