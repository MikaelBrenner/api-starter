package db

import (
	"api-starter/utils"
	"log"

	"github.com/goonode/mogo"
)

var mongoConnection *mogo.Connection = nil


func GetConnection() *mogo.Connection {
	if mongoConnection == nil {
		connectionString := utils.EnvVar("DB_CONNECTION_STRING")
		dbName := utils.EnvVar("DB_NAME")
		config := &mogo.Config{
			ConnectionString: connectionString,
			Database:         dbName,
		}
		mongoConnection, err := mogo.Connect(config)
		if err != nil {
			log.Fatal(err)
		} else {
			return mongoConnection
		}
	}
	return mongoConnection
}
