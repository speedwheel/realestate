package models

import "github.com/globalsign/mgo/bson"

type Property struct {
	ID    bson.ObjectId `bson:"_id"`
	Name  string        `bson:"name"`
	Rooms int8          `bson:"rooms"`
	Baths int8          `bson:"baths"`
	Price float64       `bson:"price"`
}
