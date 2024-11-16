package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Pelanggan struct {
	ID             primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Kode_Pelanggan string             `json:"kode_pelanggan" bson:"kode_pelanggan"`
	Tipe_Pelanggan string             `json:"tipe_pelanggan" bson:"tipe_pelanggan"`
	Nama_Pelanggan string             `json:"nama_pelanggan,omitempty" bson:"nama_pelanggan,omitempty"`
}
