package db

import (
	"fmt"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gorm_logger "logger/gorm"
	"strconv"
	"test_container"
	"time"
)

type MysqlContainer struct {
	pool *dockertest.Pool
	opts test_container.Opts
}

func NewMysqlContainer() MysqlContainer {
	opts := test_container.Opts{
		User:     "forum",
		Password: "secret",
		Database: "gforum",
		Port:     3306,
	}
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatal().Err(err).Msg("Could not connect to docker")
	}
	pool.MaxWait = time.Minute * 2

	return MysqlContainer{
		pool: pool,
		opts: opts,
	}

}
func (c MysqlContainer) CreateContainer() *dockertest.Resource {
	portInStr := strconv.Itoa(c.opts.Port)
	runOptions := &dockertest.RunOptions{
		Repository: "mysql",
		Tag:        "latest",
		Env: []string{
			"MYSQL_ROOT_PASSWORD=" + c.opts.Password,
			"MYSQL_USER=" + c.opts.User,
			"MYSQL_PASSWORD=" + c.opts.Password,
			"MYSQL_DATABASE=" + c.opts.Database,
		},
		ExposedPorts: []string{portInStr},
		Cmd: []string{
			"mysqld",
			"--character-set-server=utf8mb4",
			"--collation-server=utf8mb4_unicode_ci",
		},
		PortBindings: map[docker.Port][]docker.PortBinding{
			docker.Port(portInStr): {{HostIP: "0.0.0.0", HostPort: portInStr}},
		},
	}

	resource, err := c.pool.RunWithOptions(runOptions)

	if err != nil {
		log.Fatal().Err(err).Msgf("Could not start resource (Mysql Test Container): %s", err)
	}

	return resource
}

func (c MysqlContainer) CloseContainer(resource *dockertest.Resource) {
	if err := c.pool.Purge(resource); err != nil {
		log.Fatal().Err(err).Msgf("Could not purge resource: %s", err)
	}
}

func (c MysqlContainer) DbManager() *gorm.DB {
	var db *gorm.DB
	defaultDsn := "%s:%s@tcp(localhost:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(defaultDsn, c.opts.User, c.opts.Password, c.opts.Port, c.opts.Database)
	log.Info().Msgf("Connecting to mysql: %s", dsn)

	if err := c.pool.Retry(func() error {
		var err error
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: gorm_logger.New(),
		})
		if err != nil {
			log.Error().Err(err).Msgf("Failed to Open database: %s", err)
		}
		return nil
	}); err != nil {
		log.Fatal().Err(err).Msgf("Could not connect to docker: %s", err)
	}

	return db
}
