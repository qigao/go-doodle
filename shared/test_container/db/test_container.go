package db

import (
	"github.com/ory/dockertest/v3"
	"gorm.io/gorm"
)

type IContainer interface {
	CreateContainer() *dockertest.Resource
	DbManager() *gorm.DB
	CloseContainer(*dockertest.Resource)
}
