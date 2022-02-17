package db

import (
	"containers"
	"context"
	"github.com/rs/zerolog/log"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"time"
)

type MySQLContainer struct {
	userName *string
	password *string
	dbName   *string
	mysqlC   testcontainers.Container
	ctx      *context.Context
}

const mysqlPort = "3306/tcp"

func NewMysqlContainer(user, password, dbname string) *MySQLContainer {
	return &MySQLContainer{
		userName: &user,
		password: &password,
		dbName:   &dbname,
	}
}

func (m *MySQLContainer) CreateContainer() {
	log.Info().Msg("setup MySQL Container")
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        containers.MysqlImage,
		ExposedPorts: []string{mysqlPort},
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": *m.password,
			"MYSQL_USER":          *m.userName,
			"MYSQL_PASSWORD":      *m.password,
			"MYSQL_DATABASE":      *m.dbName,
		},
		BindMounts: map[string]string{
			"my.cnf": "/etc/mysql/my.cnf",
		},
		WaitingFor: wait.ForLog("port: 3306  MySQL Community Server - GPL").WithStartupTimeout(time.Minute * 2),
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

func (m *MySQLContainer) GetConnHostAndPort() (string, int, error) {
	host, err := m.mysqlC.Host(*m.ctx)
	if err != nil {
		return "", 0, err
	}
	p, err := m.mysqlC.MappedPort(*m.ctx, mysqlPort)
	if err != nil {
		return "", 0, err
	}
	return host, p.Int(), nil
}

func (m *MySQLContainer) CloseContainer() {
	log.Info().Msg("terminating container")
	err := m.mysqlC.Terminate(*m.ctx)
	if err != nil {
		log.Fatal().Msgf("error terminating mysql container: %s", err)
	}
}
