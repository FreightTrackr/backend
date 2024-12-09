package utils

import (
	"time"

	"github.com/FreightTrackr/backend/helpers"
	"github.com/FreightTrackr/backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertTransaksi(mongoenv *mongo.Database, collname string, datatransaksi models.Transaksi) (interface{}, error) {
	return helpers.InsertOneDoc(mongoenv, collname, datatransaksi)
}

func GetAllTransaksi(mongoenv *mongo.Database, collname string) ([]models.Transaksi, error) {
	return helpers.GetAllDoc[models.Transaksi](mongoenv, collname)
}

func GetAllTransaksiWithPagination(mongoenv *mongo.Database, collname, kode_pelanggan, no_pend string, page, limit int, startDate, endDate time.Time) ([]models.Transaksi, models.DataCount, error) {
	return helpers.GetDataForDashboard[models.Transaksi](mongoenv, collname, kode_pelanggan, no_pend, page, limit, startDate, endDate)
}

func GetAllTransaksiForVisualization(mongoenv *mongo.Database, collname string, startDate, endDate time.Time) ([]models.Transaksi, error) {
	filter := bson.M{
		"tanggal_kirim": bson.M{"$gte": startDate, "$lte": endDate},
	}
	return helpers.GetAllDocByFilter[models.Transaksi](mongoenv, collname, filter)
}

func FindTransaksi(mongoenv *mongo.Database, collname string, datatransaksi models.Transaksi) models.Transaksi {
	filter := bson.M{"no_resi": datatransaksi.No_Resi}
	return helpers.GetOneDoc[models.Transaksi](mongoenv, collname, filter)
}

func TransaksiExists(mongoenv *mongo.Database, collname string, datatransaksi models.Transaksi) bool {
	filter := bson.M{"no_resi": datatransaksi.No_Resi}
	return helpers.DocExists[models.Transaksi](mongoenv, collname, filter, datatransaksi)
}

func UpdateTransaksi(mongoenv *mongo.Database, collname string, datatransaksi models.Transaksi) interface{} {
	filter := bson.M{"no_resi": datatransaksi.No_Resi}
	return helpers.ReplaceOneDoc(mongoenv, collname, filter, datatransaksi)
}

func DeleteTransaksi(mongoenv *mongo.Database, collname string, datatransaksi models.Transaksi) interface{} {
	filter := bson.M{"no_resi": datatransaksi.No_Resi}
	return helpers.DeleteOneDoc(mongoenv, collname, filter)
}
