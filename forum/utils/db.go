package utils

import (
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	gorm_logger "logger/gorm"
	"os"
	"time"
)

func NewSqliteORM() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("../realworld.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("storage err: ", err)
	}
	return db
}

func NewMysqlORM() (*gorm.DB, error) {
	s, err := dsn()
	if err != nil {
		return nil, err
	}

	var d *gorm.DB
	for i := 0; i < 10; i++ {
		d, err = gorm.Open(mysql.Open(s), &gorm.Config{Logger: gorm_logger.New()})
		if err == nil {
			break
		}
		time.Sleep(1 * time.Second)
	}

	if err != nil {
		return nil, err
	}

	err = d.Use(
		dbresolver.Register(dbresolver.Config{ /* xxx */ }).
			SetConnMaxIdleTime(time.Hour).
			SetConnMaxLifetime(24 * time.Hour).
			SetMaxIdleConns(100).
			SetMaxOpenConns(200),
	)
	if err != nil {
		return nil, err
	}

	return d, nil
}
func NewMysqlManager() (*sql.DB, error) {
	s, err := dsn()
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("mysql", s)
	if err != nil {
		return nil, err
	}
	return db, nil
}
func dsn() (string, error) {
	host := os.Getenv("DB_HOST")
	if host == "" {
		return "", errors.New("$DB_HOST is not set")
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		return "", errors.New("$DB_USER is not set")
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		return "", errors.New("$DB_PASSWORD is not set")
	}

	name := os.Getenv("DB_NAME")
	if name == "" {
		return "", errors.New("$DB_NAME is not set")
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		return "", errors.New("$DB_PORT is not set")
	}

	options := "charset=utf8mb4&parseTime=True&loc=Local"

	// "user:password@host:port/dbname?option1&option2"
	return fmt.Sprintf("%s:%s@(%s:%s)/%s?%s",
		user, password, host, port, name, options), nil
}
