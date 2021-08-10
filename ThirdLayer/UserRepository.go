package Repository

import (
	"log"
	"context"
	//"GO_REST_API/UserStruct"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"fmt"
)

func FindUser(id string, c *gin.Context, collection *mongo.Collection){

	cur, err := collection.Find(context.TODO(), bson.D{{}})

	if err != nil {
		log.Fatal(err) }

	//var users Helper.Users
	//var user Helper.User

	for cur.Next(context.TODO()) {
		
		elem := &bson.D{}

		if err := cur.Decode(elem); err != nil {
				log.Fatal(err)
		}

		fmt.Println(elem)
		fmt.Println()

		//users = append(users, user)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
}

func FindUsers(c *gin.Context, collection *mongo.Collection){

}