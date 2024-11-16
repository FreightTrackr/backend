package utils

import (
	"github.com/FreightTrackr/backend/helpers"
	"github.com/FreightTrackr/backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertKantor(mongoenv *mongo.Database, collname string, datakantor models.Kantor) (interface{}, error) {
	return helpers.InsertOneDoc(mongoenv, collname, datakantor)
}

func GetAllKantor(mongoenv *mongo.Database, collname string) ([]models.Kantor, error) {
	return helpers.GetAllDoc[models.Kantor](mongoenv, collname)
}

func GetAllKantorWithPagination(mongoenv *mongo.Database, collname string, page, limit int) ([]models.Kantor, models.DataCount, error) {
	filter := bson.M{}
	return helpers.GetAllDocByFilterWithPagination[models.Kantor](mongoenv, collname, page, limit, filter)
}

func FindKantor(mongoenv *mongo.Database, collname string, datakantor models.Kantor) models.Kantor {
	filter := bson.M{"no_pend": datakantor.No_Pend}
	return helpers.GetOneDoc[models.Kantor](mongoenv, collname, filter)
}

func KantorExists(mongoenv *mongo.Database, collname string, datakantor models.Kantor) bool {
	filter := bson.M{"no_pend": datakantor.No_Pend}
	return helpers.DocExists[models.Kantor](mongoenv, collname, filter, datakantor)
}

func UpdateKantor(mongoenv *mongo.Database, collname string, datakantor models.Kantor) interface{} {
	filter := bson.M{"no_pend": datakantor.No_Pend}
	return helpers.ReplaceOneDoc(mongoenv, collname, filter, datakantor)
}

func DeleteKantor(mongoenv *mongo.Database, collname string, datakantor models.Kantor) interface{} {
	filter := bson.M{"no_pend": datakantor.No_Pend}
	return helpers.DeleteOneDoc(mongoenv, collname, filter)
}
