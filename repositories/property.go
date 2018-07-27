package repositories

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/speedwheel/immigrationrealestate/datasource/mongo"
	"github.com/speedwheel/immigrationrealestate/models"
)

const (
	propertyCollectionName = "properties"
)

// SelectPropertyByQUery selects all the properties from the database that
// match with the given query
// It accepts a database connection pointer, a bson.M containing the query information
// and a bson.M defining the selected fields that need to be returned
// It returns a slice of properties and an error
func SelectPropertiesByQUery(db *mongo.MgoDb, query bson.M, selectedFields bson.M) (properties []models.Property, err error) {
	c := db.C(propertyCollectionName)
	err = c.Find(query).Select(selectedFields).All(&properties)
	return
}

// SelectPropertyByQuery selects selects a single property from the database that
// match with the given query
// It accepts a database connection pointer, a bson.M containing the query information
// and a bson.M defining the selected fields that need to be returned
// It returns a single property and an error
func SelectPropertyByQuery(db *mongo.MgoDb, query bson.M, selectedFields bson.M) (property models.Property, err error) {
	c := db.C(propertyCollectionName)
	err = c.Find(query).Select(selectedFields).One(&property)
	return
}

// InsertProperties inserts one or more properties in the database,
// It accepts a database connection pointer and one or more interfaces
// of type models.Property
// It returns an error
func InsertProperties(db *mongo.MgoDb, properties ...interface{}) (err error) {
	c := db.C(propertyCollectionName)
	err = c.Insert(properties...)
	return
}

// UpdateProperty updates a single property
// It accespts a database connection pointer, a bson.M with the query
// conditions, and another bson.M with the new updated values
// It returns an error
func UpdateProperty(db *mongo.MgoDb, query bson.M, update bson.M) (err error) {
	c := db.C(propertyCollectionName)
	err = c.Update(query, update)
	return
}

// UpdateProperties updates one or more properties
// It accespts a database connection pointer, a bson.M with the query
// conditions, and another bson.M with the new updated values
// It returns a struct with nunumber of existing documents modified and
// and error
func UpdateProperties(db *mongo.MgoDb, query bson.M, update bson.M) (info *mgo.ChangeInfo, err error) {
	c := db.C(propertyCollectionName)
	info, err = c.UpdateAll(query, update)
	return
}

// RemoveProperty removes a single property from the database
// It accespts a database connection pointer and a bson.M with the query
// It returns an error
func RemoveProperty(db *mongo.MgoDb, query bson.M) (err error) {
	c := db.C(propertyCollectionName)
	err = c.Remove(query)
	return
}

// RemoveProperties removes one or more properties from the database
// It accespts a database connection pointer and a bson.M with the query
// It returns a struct with nunumber of existing documents modified and
// and error
func RemoveProperties(db *mongo.MgoDb, query bson.M) (info *mgo.ChangeInfo, err error) {
	c := db.C(propertyCollectionName)
	info, err = c.RemoveAll(query)
	return
}
