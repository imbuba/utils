package mysql

import (
	"os"

	"github.com/imbuba/utils/database"
	"github.com/jinzhu/gorm"

	// needs
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB link to DB
var DB *gorm.DB

var (
	config *database.Database
	dirs   = []string{"conf/", ""}
	pathes = []string{"mysql.config", "db.config", "mysql_config.json", "db_config.json"}
)

func init() {
	var err error
	dbConfig := os.Getenv("MYSQL_CONFIG")
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
		config.Connection = "%{User}%:%{Password}%@%{Protocol}%(%{Host}%:%{Port}%)/%{DBName}%?charset=utf8&parseTime=True&loc=Local"
	}
	DB, err = gorm.Open("mysql", config.ConnectionString())
	if err != nil {
		panic(err)
	}
}
