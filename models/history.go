package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Lokasi struct {
	Status     string    `json:"status" bson:"status"`
	Timestamp  time.Time `json:"timestamp" bson:"timestamp"`
	Coordinate []float64 `json:"coordinate" bson:"coordinate"`
	Catatan    string    `json:"catatan" bson:"catatan"`
	Username   string    `json:"username" bson:"username"`
}

type History struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	ID_History string             `json:"id_history" bson:"id_history"`
	Lokasi     []Lokasi           `json:"lokasi" bson:"lokasi"`
}
