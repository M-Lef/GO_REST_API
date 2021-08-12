package Repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/gin-gonic/gin"
	"GO_REST_API/User_Data"
	"log"
	"context"
)

func PostUser(new_users []User.User, c *gin.Context, collection *mongo.Collection) {
	for _ , a := range new_users {
		_, err := collection.InsertOne(context.Background(), a)
		if err != nil {
			log.Fatal(err)
		}
	}
}