package mssql

import (
	"os"

	"github.com/imbuba/utils/database"
	"github.com/jinzhu/gorm"

	// needs
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

// DB link to DB
var DB *gorm.DB

var (
	config *database.Database
	dirs   = []string{"conf/", ""}
	pathes = []string{"mssql.config", "db.config", "mssql_config.json", "db_config.json"}
)

func init() {
	var err error
	dbConfig := os.Getenv("MSSQL_CONFIG")
	if dbConfig != "" {
		config, _ = database.ParseConfig(dbConfig)
	}
	if config == nil {
	out:
		for _, d := range dirs {
			for _, p := range pathes {
				if config, err = database.ParseConfig(d + p); err == nil {
					break out
				}
			}
		}
	}
	if config == nil {
		panic("Config not found")
	}
	if config.Connection == "" {
		config.Connection = "user id=%{User}%;password=%{Password}%;server=%{Host}%;port=%{Port}%;database=%{DBName}%;app name=%{AppName}%;connection timeout=36000"
	}
	DB, err = gorm.Open("mssql", config.ConnectionString())
	if err != nil {
		panic(err)
	}
}
