package utils

import (
	"github.com/FreightTrackr/backend/helpers"
	"github.com/FreightTrackr/backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertHistory(mongoenv *mongo.Database, collname string, datahistory models.History) (interface{}, error) {
	return helpers.InsertOneDoc(mongoenv, collname, datahistory)
}

func GetAllHistory(mongoenv *mongo.Database, collname string) ([]models.History, error) {
	return helpers.GetAllDoc[models.History](mongoenv, collname)
}

func GetAllHistoryWithPagination(mongoenv *mongo.Database, collname string, page, limit int) ([]models.History, models.DataCount, error) {
	return helpers.GetAllDocWithPagination[models.History](mongoenv, collname, page, limit)
}

func FindHistory(mongoenv *mongo.Database, collname string, datahistory models.History) models.History {
	filter := bson.M{"id_history": datahistory.ID_History}
	return helpers.GetOneDoc[models.History](mongoenv, collname, filter)
}

func HistoryExists(mongoenv *mongo.Database, collname string, datahistory models.History) bool {
	filter := bson.M{"id_history": datahistory.ID_History}
	return helpers.DocExists[models.History](mongoenv, collname, filter, datahistory)
}

func UpdateHistory(mongoenv *mongo.Database, collname string, datahistory models.History) interface{} {
	filter := bson.M{"id_history": datahistory.ID_History}
	return helpers.ReplaceOneDoc(mongoenv, collname, filter, datahistory)
}

func DeleteHistory(mongoenv *mongo.Database, collname string, datahistory models.History) interface{} {
	filter := bson.M{"id_history": datahistory.ID_History}
	return helpers.DeleteOneDoc(mongoenv, collname, filter)
}
