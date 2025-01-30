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

func GetAllTransaksiWithPagination(mongoenv *mongo.Database, collname, kode_pelanggan, no_pend string, page, limit int, startDate, endDate time.Time) ([]models.Transaksi, models.DataCount, error) {
	return helpers.GetDataForDashboard[models.Transaksi](mongoenv, collname, kode_pelanggan, no_pend, page, limit, startDate, endDate)
}

func GetAllTransaksi(mongoenv *mongo.Database, collname, no_pend, kode_pelanggan string, startDate, endDate time.Time) ([]models.Transaksi, error) {
	filter := bson.M{
		"tanggal_kirim": bson.M{"$gte": startDate, "$lte": endDate},
	}
	if no_pend != "" {
		filter["$or"] = []bson.M{
			{"no_pend_kirim": no_pend},
			{"no_pend_terima": no_pend},
		}
	}
	if kode_pelanggan != "" {
		filter["kode_pelanggan"] = kode_pelanggan
	}
	return helpers.GetAllDocByFilter[models.Transaksi](mongoenv, collname, filter)
}

func GetTransaksiTesting(mongoenv *mongo.Database, collname string, limit int, startDate, endDate time.Time) ([]models.Transaksi, error) {
	filter := bson.M{
		"tanggal_kirim": bson.M{"$gte": startDate, "$lte": endDate},
	}
	return helpers.GetDocTesting[models.Transaksi](mongoenv, collname, limit, filter)
}

func GetStatusDeliveredTransaksi(mongoenv *mongo.Database, collname, no_pend, kode_pelanggan string, startDate, endDate time.Time) ([]models.Transaksi, error) {
	filter := bson.M{
		"tanggal_kirim": bson.M{"$gte": startDate, "$lte": endDate},
		"status":        "delivered",
	}
	if no_pend != "" {
		filter["$or"] = []bson.M{
			{"no_pend_kirim": no_pend},
			{"no_pend_terima": no_pend},
		}
	}
	if kode_pelanggan != "" {
		filter["kode_pelanggan"] = kode_pelanggan
	}
	return helpers.GetAllDocByFilter[models.Transaksi](mongoenv, collname, filter)
}

func GetTipeCodTransaksi(mongoenv *mongo.Database, collname, no_pend, kode_pelanggan string, startDate, endDate time.Time) ([]models.Transaksi, error) {
	filter := bson.M{
		"tanggal_kirim": bson.M{"$gte": startDate, "$lte": endDate},
		"status":        "delivered",
		"tipe_cod":      "cod",
	}
	if no_pend != "" {
		filter["$or"] = []bson.M{
			{"no_pend_kirim": no_pend},
			{"no_pend_terima": no_pend},
		}
	}
	if kode_pelanggan != "" {
		filter["kode_pelanggan"] = kode_pelanggan
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
