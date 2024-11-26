package helpers

import (
	"context"
	"fmt"
	"time"

	"github.com/FreightTrackr/backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnect(mconn models.DBInfo) (db *mongo.Database) {
	clientOptions := options.Client().ApplyURI((mconn.DBString))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Printf("Error connecting to MongoDB: %v", err)
	}
	return client.Database(mconn.DBName)
}

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}, err error) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
		return nil, err
	}
	return insertResult.InsertedID, nil
}

func GetOneDoc[T any](db *mongo.Database, collection string, filter bson.M) (doc T) {
	err := db.Collection(collection).FindOne(context.TODO(), filter).Decode(&doc)
	if err != nil {
		fmt.Printf("GetOneDoc: %v\n", err)
	}
	return
}

func GetOneLatestDoc[T any](db *mongo.Database, collection string, filter bson.M) (doc T, err error) {
	opts := options.FindOne().SetSort(bson.M{"$natural": -1})
	err = db.Collection(collection).FindOne(context.TODO(), filter, opts).Decode(&doc)
	return doc, err
}

func GetAllDocByFilter[T any](db *mongo.Database, collection string, filter bson.M) (doc []T, err error) {
	ctx := context.TODO()
	cur, err := db.Collection(collection).Find(ctx, filter)
	if err != nil {
		fmt.Printf("GetAllDoc: %v\n", err)
		return nil, err
	}
	defer cur.Close(ctx)
	err = cur.All(ctx, &doc)
	if err != nil {
		fmt.Printf("GetAllDoc Cursor Err: %v\n", err)
		return nil, err
	}
	return doc, nil
}
func GetAllDocByFilterWithPagination[T any](db *mongo.Database, collection string, page, limit int, filter bson.M) (doc []T, datacount models.DataCount, err error) {
	ctx := context.TODO()
	findOptions := options.Find()
	findOptions.SetSkip(int64((page - 1) * limit))
	findOptions.SetLimit(int64(limit))
	cur, err := db.Collection(collection).Find(ctx, filter, findOptions)
	if err != nil {
		fmt.Printf("GetAllDoc: %v\n", err)
		return nil, models.DataCount{Total: 0}, err
	}
	defer cur.Close(ctx)
	err = cur.All(ctx, &doc)
	if err != nil {
		fmt.Printf("GetAllDoc Cursor Err: %v\n", err)
		return nil, models.DataCount{Total: 0}, err
	}
	count, err := db.Collection(collection).CountDocuments(ctx, filter)
	if err != nil {
		fmt.Printf("GetAllDoc CountDocuments Error: %v\n", err)
		return nil, models.DataCount{Total: 0}, err
	}
	return doc, models.DataCount{Total: count}, nil
}

func GetAllDoc[T any](db *mongo.Database, collection string) (doc []T, err error) {
	ctx := context.TODO()
	cur, err := db.Collection(collection).Find(ctx, bson.M{})
	if err != nil {
		fmt.Printf("GetAllDoc: %v\n", err)
		return nil, err
	}
	defer cur.Close(ctx)
	err = cur.All(ctx, &doc)
	if err != nil {
		fmt.Printf("GetAllDoc Cursor Err: %v\n", err)
		return nil, err
	}
	return doc, nil
}

func GetAllDocWithPagination[T any](db *mongo.Database, collection string, page, limit int) (doc []T, datacount models.DataCount, err error) {
	ctx := context.TODO()
	findOptions := options.Find()
	findOptions.SetSkip(int64((page - 1) * limit))
	findOptions.SetLimit(int64(limit))
	cur, err := db.Collection(collection).Find(ctx, bson.M{}, findOptions)
	if err != nil {
		fmt.Printf("GetAllDoc: %v\n", err)
		return nil, models.DataCount{Total: 0}, err
	}
	defer cur.Close(ctx)
	err = cur.All(ctx, &doc)
	if err != nil {
		fmt.Printf("GetAllDoc Cursor Err: %v\n", err)
		return nil, models.DataCount{Total: 0}, err
	}
	count, err := db.Collection(collection).CountDocuments(ctx, bson.M{})
	if err != nil {
		fmt.Printf("GetAllDoc CountDocuments Error: %v\n", err)
		return nil, models.DataCount{Total: 0}, err
	}
	return doc, models.DataCount{Total: count}, nil
}

func GetDataForDashboard[T any](db *mongo.Database, collection, kode_pelanggan, no_pend string, page, limit int, startDate, endDate time.Time) (docs []T, datacount models.DataCount, err error) {
	ctx := context.TODO()
	findOptions := options.Find()
	findOptions.SetSkip(int64((page - 1) * limit))
	findOptions.SetLimit(int64(limit))
	statuses := []string{"delivered", "canceled", "returned", "inWarehouse", "inVehicle", "failed"}

	// Tambahkan filter untuk tanggal_kirim dan tanggal_terima
	filter := bson.M{
		"tanggal_kirim": bson.M{"$gte": startDate, "$lte": endDate},
	}

	if kode_pelanggan != "" {
		filter["kode_pelanggan"] = kode_pelanggan
	}

	if no_pend != "" {
		filter["$or"] = []bson.M{
			{"no_pend_kirim": no_pend},
			{"no_pend_terima": no_pend},
		}
	}

	cur, err := db.Collection(collection).Find(ctx, filter, findOptions)
	if err != nil {
		fmt.Printf("GetAllDoc Find Error: %v\n", err)
		return nil, models.DataCount{}, err
	}
	defer cur.Close(ctx)

	err = cur.All(ctx, &docs)
	if err != nil {
		fmt.Printf("GetAllDoc Cursor Error: %v\n", err)
		return nil, models.DataCount{}, err
	}

	datacount.Total, err = db.Collection(collection).CountDocuments(ctx, filter)
	if err != nil {
		fmt.Printf("GetAllDoc CountDocuments Error: %v\n", err)
		return nil, models.DataCount{}, err
	}

	for _, status := range statuses {
		statusFilter := bson.M{
			"status":        status,
			"tanggal_kirim": bson.M{"$gte": startDate, "$lte": endDate},
		}
		if kode_pelanggan != "" {
			statusFilter["kode_pelanggan"] = kode_pelanggan
		}
		if no_pend != "" {
			statusFilter["$or"] = []bson.M{
				{"no_pend_kirim": no_pend},
				{"no_pend_terima": no_pend},
			}
		}
		count, err := db.Collection(collection).CountDocuments(ctx, statusFilter)
		if err != nil {
			fmt.Printf("GetAllDoc CountDocuments for %s Error: %v\n", status, err)
			return nil, models.DataCount{}, err
		}
		switch status {
		case "delivered":
			datacount.Delivered = count
		case "canceled":
			datacount.Canceled = count
		case "returned":
			datacount.Returned = count
		case "inWarehouse":
			datacount.InWarehouse = count
		case "inVehicle":
			datacount.InVehicle = count
		case "failed":
			datacount.Failed = count
		}
	}
	return docs, datacount, nil
}

func GetAllDistinctDoc(db *mongo.Database, filter bson.M, fieldname, collection string) (doc []any) {
	ctx := context.TODO()
	doc, err := db.Collection(collection).Distinct(ctx, fieldname, filter)
	if err != nil {
		fmt.Printf("GetAllDistinctDoc: %v\n", err)
	}
	return
}

func ReplaceOneDoc(db *mongo.Database, collection string, filter bson.M, doc interface{}) (updatereseult *mongo.UpdateResult) {
	updatereseult, err := db.Collection(collection).ReplaceOne(context.TODO(), filter, doc)
	if err != nil {
		fmt.Printf("ReplaceOneDoc: %v\n", err)
	}
	return
}

func DeleteOneDoc(db *mongo.Database, collection string, filter bson.M) (result *mongo.DeleteResult) {
	result, err := db.Collection(collection).DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Printf("DeleteOneDoc: %v\n", err)
	}
	return
}

func DeleteDoc(db *mongo.Database, collection string, filter bson.M) (result *mongo.DeleteResult) {
	result, err := db.Collection(collection).DeleteMany(context.TODO(), filter)
	if err != nil {
		fmt.Printf("DeleteDoc : %v\n", err)
	}
	return
}

func GetRandomDoc[T any](db *mongo.Database, collection string, size uint) (result []T, err error) {
	filter := mongo.Pipeline{
		{{Key: "$sample", Value: bson.D{{Key: "size", Value: size}}}},
	}
	ctx := context.Background()
	cur, err := db.Collection(collection).Aggregate(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	if err = cur.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func DocExists[T any](db *mongo.Database, collname string, filter bson.M, doc T) (result bool) {
	err := db.Collection(collname).FindOne(context.Background(), filter).Decode(&doc)
	return err == nil
}
