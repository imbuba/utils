package mongo

import (
	"log"
	"os"

	"github.com/imbuba/utils/database"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// DBWrapper struct
type DBWrapper struct {
	config  *database.Database
	session *mgo.Session
}

// DB link to DB
var DB = &DBWrapper{}

var (
	config *database.Database
	dirs   = []string{"conf/", ""}
	pathes = []string{"mongo.config", "db.config", "mongo_config.json", "db_config.json"}
)

func init() {
	var err error
	dbConfig := os.Getenv("MONGO_CONFIG")
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
		log.Println("Config not found")
		return
	}
	DB.config = config
	if config.Connection == "" {
		config.Connection = "mongodb://"
		for i, h := range config.Hosts {
			config.Connection += h + ":" + config.Ports[i]
		}
		config.Connection += "/" + config.DBName
	}
	DB.session, err = mgo.Dial(config.Connection)
	if err != nil {
		panic(err)
	}
	// mgo.SetDebug(true)
	// var aLogger *log.Logger
	// aLogger = log.New(os.Stderr, "", log.LstdFlags)
	// mgo.SetLogger(aLogger)
}

// Get returns database and close function
func (d *DBWrapper) Get() (*mgo.Database, func()) {
	newSess := d.session.Copy()
	return newSess.DB(d.config.DBName), newSess.Close
}

// Custom returns database with custom name and close function
func (d *DBWrapper) Custom(name string) (*mgo.Database, func()) {
	newSess := d.session.Copy()
	return newSess.DB(name), newSess.Close
}

// SetPoolLimit sets pool limit
func (d *DBWrapper) SetPoolLimit(limit int) {
	d.session.SetPoolLimit(limit)
}

type autoincrDoc struct {
	N int `bson:"n"`
}

// GetNewID returns new autoincrement id
func GetNewID(collName string) (int, error) {
	db, closeDB := DB.Get()
	defer closeDB()
	doc := autoincrDoc{}
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"n": 1}},
		Upsert:    true,
		ReturnNew: true,
	}
	_, err := db.C("autoincrements").Find(bson.M{"_id": collName}).Apply(change, &doc)
	if err != nil {
		return 0, err
	}
	return doc.N, nil
}
