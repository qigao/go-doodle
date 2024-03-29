package main

import (
	"db"
	"flag"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	pathPtr := flag.String("path", ".", "path of sqlboiler.toml")
	dbPtr := flag.String("db", "db/sql", "a string")
	flag.Parse()
	config := db.ReadDBConfigFromToml(*pathPtr)
	dsn := db.BuildDSNFromDbConfig(config, config.Host, config.Port)
	db := db.SqlDbManager(dsn)
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	mg, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", *dbPtr), "mysql", driver)
	if err != nil {
		panic(err.Error())
	}
	mg.Up()
}
