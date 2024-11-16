package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Kantor struct {
	ID              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	No_Pend         string             `json:"no_pend" bson:"no_pend"`
	No_Pend_Kcu     string             `json:"no_pend_kcu,omitempty" bson:"no_pend_kcu,omitempty"`
	No_Pend_Kc      string             `json:"no_pend_kc,omitempty" bson:"no_pend_kc,omitempty"`
	Tipe_Kantor     string             `json:"tipe_kantor" bson:"tipe_kantor"`
	Nama_Kantor     string             `json:"nama_kantor" bson:"nama_kantor"`
	Region_Kantor   int                `json:"region_kantor" bson:"region_kantor"`
	Kota_Kantor     string             `json:"kota_kantor" bson:"kota_kantor"`
	Kode_Pos_Kantor int                `json:"kode_pos_kantor" bson:"kode_pos_kantor"`
	Alamat_Kantor   string             `json:"alamat_kantor" bson:"alamat_kantor"`
}
