package routes

import (
	"github.com/kataras/iris"
	"github.com/speedwheel/immigrationrealestate/datasource/mongo"
	"github.com/speedwheel/immigrationrealestate/repositories"
)

func GetHomeHandler(ctx iris.Context, db *mongo.MgoDb) {
	/*propertiesToInsert := []interface{}{
		models.Property{
			ID:   bson.NewObjectId(),
			Name: "1 Bedroom Apartment",
		},
		models.Property{
			ID:   bson.NewObjectId(),
			Name: "5 Bedroom Apartment",
		},
	}
	err := repositories.InsertProperties(db, propertiesToInsert...)
	fmt.Println(err)*/
	properties, _ := repositories.SelectPropertiesByQUery(db, nil, nil)
	ctx.JSON(properties)
}
