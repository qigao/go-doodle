package db

import (
	"containers"
	"context"
	"os"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
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
	ctx := context.Background()
	seedDataPath, _ := os.Getwd()
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
			seedDataPath + "/my.cnf": "/etc/mysql/my.cnf",
		},
		WaitingFor: wait.ForLog("port: 3306  MySQL Community Server").WithStartupTimeout(time.Minute * 2),
	}
	mysqlC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	if err != nil {
		panic("error getting mysql container: " + err.Error())
	}
	m.mysqlC = mysqlC
	m.ctx = &ctx
}

func (m *MySQLContainer) GetConnHostAndPort() (string, int, error) {
	host, err := m.mysqlC.Host(*m.ctx)
	if err != nil {
		panic("error getting host:" + err.Error())
	}
	p, err := m.mysqlC.MappedPort(*m.ctx, mysqlPort)
	if err != nil {
		panic("error getting port:" + err.Error())
	}
	return host, p.Int(), nil
}

func (m *MySQLContainer) CloseContainer() {
	err := m.mysqlC.Terminate(*m.ctx)
	if err != nil {
		panic("error terminating container:" + err.Error())
	}
}
