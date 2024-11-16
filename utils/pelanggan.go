package utils

import (
	"github.com/FreightTrackr/backend/helpers"
	"github.com/FreightTrackr/backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertPelanggan(mongoenv *mongo.Database, collname string, datapelanggan models.Pelanggan) (interface{}, error) {
	return helpers.InsertOneDoc(mongoenv, collname, datapelanggan)
}

func GetAllPelanggan(mongoenv *mongo.Database, collname string) ([]models.Pelanggan, error) {
	return helpers.GetAllDoc[models.Pelanggan](mongoenv, collname)
}

func GetAllPelangganWithPagination(mongoenv *mongo.Database, collname string, page, limit int) ([]models.Pelanggan, models.DataCount, error) {
	return helpers.GetAllDocWithPagination[models.Pelanggan](mongoenv, collname, page, limit)
}

func GetAllPelangganByFilterWithPagination(mongoenv *mongo.Database, collname string, page, limit int, tipe_pelanggan string) ([]models.Pelanggan, models.DataCount, error) {
	filter := bson.M{"tipe_pelanggan": tipe_pelanggan}
	return helpers.GetAllDocByFilterWithPagination[models.Pelanggan](mongoenv, collname, page, limit, filter)
}

func FindPelanggan(mongoenv *mongo.Database, collname string, datapelanggan models.Pelanggan) models.Pelanggan {
	filter := bson.M{"kode_pelanggan": datapelanggan.Kode_Pelanggan}
	return helpers.GetOneDoc[models.Pelanggan](mongoenv, collname, filter)
}

func PelangganExists(mongoenv *mongo.Database, collname string, datapelanggan models.Pelanggan) bool {
	filter := bson.M{"kode_pelanggan": datapelanggan.Kode_Pelanggan}
	return helpers.DocExists[models.Pelanggan](mongoenv, collname, filter, datapelanggan)
}

func UpdatePelanggan(mongoenv *mongo.Database, collname string, datapelanggan models.Pelanggan) interface{} {
	filter := bson.M{"kode_pelanggan": datapelanggan.Kode_Pelanggan}
	return helpers.ReplaceOneDoc(mongoenv, collname, filter, datapelanggan)
}

func DeletePelanggan(mongoenv *mongo.Database, collname string, datapelanggan models.Pelanggan) interface{} {
	filter := bson.M{"kode_pelanggan": datapelanggan.Kode_Pelanggan}
	return helpers.DeleteOneDoc(mongoenv, collname, filter)
}
