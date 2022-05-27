package db_tools

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type DatabaseConfig struct {
	Host   string `mapstructure:"host"`
	Port   int    `mapstructure:"port"`
	User   string `mapstructure:"user"`
	Pass   string `mapstructure:"pass"`
	DbName string `mapstructure:"dbname"`
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

	options := "charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true"

	// "user:password@host:port/dbname?option1&option2"
	return fmt.Sprintf("%s:%s@(%s:%s)/%s?%s",
		user, password, host, port, name, options), nil
}

type dbConfig struct {
	Db DatabaseConfig `mapstructure:"mysql"`
}

func BuildDSNFromDbConfig(config *DatabaseConfig, host string, port int) string {
	defaultDsn := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true"
	dsn := fmt.Sprintf(defaultDsn, config.User, config.Pass, host, port, config.DbName)
	return dsn
}

//ReadDBConfigFromToml read config from sqlboiler.toml file
func ReadDBConfigFromToml(path string) *DatabaseConfig {
	v := viper.New()
	v.SetConfigName("sqlboiler")
	v.AddConfigPath(path)
	if err := v.ReadInConfig(); err != nil {
		log.Fatal().Err(err).Msg("Failed to read config file")
	}
	var c dbConfig
	if err := v.Unmarshal(&c); err != nil {
		log.Fatal().Err(err).Msg("Failed to unmarshal config file")
	}
	return &c.Db
}

func SqlDbManager(dsn string) *sql.DB {
	var db *sql.DB
	boil.DebugMode = true
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to Open database: %s", err)
	}
	return db
}
