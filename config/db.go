package config

import (
	"os"

	"github.com/gocroot/helper/atdb"
)

var MongoString string = os.Getenv("MONGOSTRINGPM3")

var mongoinfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "PM3",
}

var Mongoconn, ErrorMongoconn = atdb.MongoConnect(mongoinfo)
