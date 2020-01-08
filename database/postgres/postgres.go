package postgres

import (
	"os"

	"github.com/imbuba/utils/database"
	"github.com/jinzhu/gorm"

	// needs
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB link to DB
var DB *gorm.DB

var (
	config *database.Database
	dirs   = []string{"conf/", ""}
	pathes = []string{"pg.config", "db.config", "pg_config.json", "db_config.json"}
)

func init() {
	var err error
	dbConfig := os.Getenv("POSTGRES_CONFIG")
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
		config.Connection = "host=%{Host}% port=%{Port}% user=%{User}% dbname=%{DBName}% password=%{Password}% sslmode=disable"
	}
	DB, err = gorm.Open("postgres", config.ConnectionString())
	if err != nil {
		panic(err)
	}
}

// AlternativeDB returns alternative db
func AlternativeDB(filename string) (*gorm.DB, error) {
	config, err := database.ParseConfig(filename)
	if err != nil {
		return nil, err
	}
	if config.Connection == "" {
		config.Connection = "host=%{Host}% port=%{Port}% user=%{User}% dbname=%{DBName}% password=%{Password}% sslmode=disable"
	}
	return gorm.Open("postgres", config.ConnectionString())
}
