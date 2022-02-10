package db

import (
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"strconv"
	"test_container"
	"time"
)

type CassandraContainer struct {
	pool     *dockertest.Pool
	resource *dockertest.Resource
	opts     test_container.Opts
}

func NewCassandraContainer() *CassandraContainer {
	opts := test_container.Opts{
		User:     "testcontainer",
		Password: "secret",
		Database: "testcontainer",
		Port:     3306,
	}
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatal().Err(err).Msg("Could not connect to docker")
	}
	pool.MaxWait = time.Minute * 2
	return &CassandraContainer{
		pool: pool,
		opts: opts,
	}
}
func (c CassandraContainer) CreateContainer() *dockertest.Resource {
	portInStr := strconv.Itoa(c.opts.Port)
	runOptions := dockertest.RunOptions{
		Repository:   "cassandra",
		Tag:          "latest",
		ExposedPorts: []string{portInStr},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"9042/tcp": {{HostIP: "0.0.0.0", HostPort: portInStr}},
		},
	}
	resource, err := c.pool.RunWithOptions(&runOptions)
	if err != nil {
		log.Fatal().Err(err).Msgf("Could not start c.resource (Cassandra Test Container): %s", err.Error())
		return nil
	}
	return resource
}

func (c CassandraContainer) DbManager() *gorm.DB {
	var db *gorm.DB
	//TODO: Add Cassandra driver for gorm

	return db
}

func (c CassandraContainer) CloseContainer(resource *dockertest.Resource) {
	if err := c.pool.Purge(resource); err != nil {
		log.Fatal().Err(err).Msgf("Could not purge c.resource: %s", err)
	}
}
