package utils

import (
	"github.com/FreightTrackr/backend/helpers"
	"github.com/FreightTrackr/backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertUser(mongoenv *mongo.Database, collname string, datauser models.Users) (interface{}, error) {
	return helpers.InsertOneDoc(mongoenv, collname, datauser)
}

func GetAllUser(mongoenv *mongo.Database, collname string) ([]models.Users, error) {
	return helpers.GetAllDoc[models.Users](mongoenv, collname)
}

func GetAllUserWithPagination(mongoenv *mongo.Database, collname string, page, limit int) ([]models.Users, models.DataCount, error) {
	filter := bson.M{}
	return helpers.GetAllDocByFilterWithPagination[models.Users](mongoenv, collname, page, limit, filter)
}

func FindUser(mongoenv *mongo.Database, collname string, datauser models.Users) models.Users {
	filter := bson.M{"username": datauser.Username}
	return helpers.GetOneDoc[models.Users](mongoenv, collname, filter)
}

func IsPasswordValid(mongoenv *mongo.Database, collname string, datauser models.Users) bool {
	filter := bson.M{"username": datauser.Username}
	res := helpers.GetOneDoc[models.Users](mongoenv, collname, filter)
	hashChecker := helpers.CheckPasswordHash(datauser.Password, res.Password)
	return hashChecker
}

func UsernameExists(mongoenv *mongo.Database, collname string, datauser models.Users) bool {
	filter := bson.M{"username": datauser.Username}
	return helpers.DocExists[models.Users](mongoenv, collname, filter, datauser)
}

func UpdateUser(mongoenv *mongo.Database, collname string, datauser models.Users) interface{} {
	filter := bson.M{"username": datauser.Username}
	return helpers.ReplaceOneDoc(mongoenv, collname, filter, datauser)
}

func DeleteUser(mongoenv *mongo.Database, collname string, datauser models.Users) interface{} {
	filter := bson.M{"username": datauser.Username}
	return helpers.DeleteOneDoc(mongoenv, collname, filter)
}
