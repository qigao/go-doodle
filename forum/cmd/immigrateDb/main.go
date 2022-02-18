package main

import (
	"flag"
	"fmt"
	"forum/utils"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	pathPtr := flag.String("path", ".", "path of sqlboiler.toml")
	dbPtr := flag.String("db", "db/sql", "a string")
	flag.Parse()
	config := utils.ReadDBConfigFromToml(*pathPtr)
	dsn := utils.BuildDSNFromDbConfig(config, config.Host, config.Port)
	fmt.Println(dsn)
	db := utils.SqlDbManager(dsn)
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	mg, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", *dbPtr), "mysql", driver)
	if err != nil {
		panic(err.Error())
	}
	mg.Up()
	//TODO: execute sqlboiler
}
