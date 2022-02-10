package db

import (
	"fmt"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
	"test_container"
	"time"
)

type PostgresSqlContainer struct {
	pool     *dockertest.Pool
	resource *dockertest.Resource
	opts     test_container.Opts
}

func NewPostgresqlContainer() PostgresSqlContainer {
	opts := test_container.Opts{
		User:     "testcontainer",
		Password: "secret",
		Database: "testcontainer",
		Port:     5432,
	}
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatal().Err(err).Msg("Could not connect to docker")
	}
	pool.MaxWait = time.Minute * 2
	return PostgresSqlContainer{opts: opts, pool: pool}
}

func (c PostgresSqlContainer) CreateContainer() *dockertest.Resource {
	portInStr := strconv.Itoa(c.opts.Port)
	dockerOpts := dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "latest",
		Env: []string{
			"POSTGRES_USER=" + c.opts.User,
			"POSTGRES_PASSWORD=" + c.opts.Password,
			"POSTGRES_DB=" + c.opts.Database,
		},
		ExposedPorts: []string{portInStr},
		PortBindings: map[docker.Port][]docker.PortBinding{
			docker.Port(portInStr): {{HostIP: "0.0.0.0", HostPort: portInStr}},
		},
	}
	resource, err := c.pool.RunWithOptions(&dockerOpts)
	if err != nil {
		log.Fatal().Err(err).Msgf("Could not start resource (Postgresql Test Container): %s", err.Error())
	}
	return resource
}

func (c PostgresSqlContainer) DbManager() *gorm.DB {
	var db *gorm.DB
	if err := c.pool.Retry(func() error {
		defaultDsn := "host=localhost user=%s password=%s dbname=%s port=%d sslmode=disable"
		dsn := fmt.Sprintf(defaultDsn, c.opts.User, c.opts.Password, c.opts.Database, c.opts.Port)

		var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal().Err(err).Msgf("Failed to Open database: %s", err)
		}

		return nil
	}); err != nil {
		log.Fatal().Err(err).Msgf("Could not connect to docker: %s", err)
	}

	return db
}

func (c PostgresSqlContainer) CloseContainer(resource *dockertest.Resource) {
	if err := c.pool.Purge(resource); err != nil {
		log.Fatal().Err(err).Msgf("Could not purge resource: %s", err)
	}
}
