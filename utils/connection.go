package utils

import (
	"os"
	"sync"

	"github.com/FreightTrackr/backend/helpers"
	"github.com/FreightTrackr/backend/models"
	"go.mongodb.org/mongo-driver/mongo"
)

var dbInstance *mongo.Database
var once sync.Once

func SetConnection() *mongo.Database {
	once.Do(func() {
		dbInfo := models.DBInfo{
			DBString: os.Getenv("DB_URI"),
			DBName:   os.Getenv("DB_DATABASE"),
		}
		dbInstance = helpers.MongoConnect(dbInfo)
	})
	if dbInstance == nil {
		return nil
	}
	return dbInstance
}
