package utils

import (
	"os"

	"github.com/FreightTrackr/backend/helpers"
	"github.com/FreightTrackr/backend/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetConnection() *mongo.Database {
	var DBmongoinfo = models.DBInfo{
		DBString: os.Getenv("DB_URI"),
		DBName:   os.Getenv("DB_DATABASE"),
	}
	return helpers.MongoConnect(DBmongoinfo)
}
