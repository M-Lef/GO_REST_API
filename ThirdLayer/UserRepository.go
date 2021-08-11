package Repository

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/gin-gonic/gin"
	"GO_REST_API/Users"
	"context"
	"log"
)

func FindUser(id string, c *gin.Context, collection *mongo.Collection) User.User{

	cur, err := collection.Find(context.TODO(), bson.D{{}})

	if err != nil {
		log.Fatal(err) 
	}

	var res User.User

	for cur.Next(context.TODO()) {
		
		var elem User.User
		
		if err := cur.Decode(&elem); err != nil {
			log.Fatal(err)
		}
		
		if (elem.ID == id) {
			return elem
		}
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return res
}

func FindUsers(c *gin.Context, collection *mongo.Collection) []*User.User{

	cur, err := collection.Find(context.TODO(), bson.D{{}})

	if err != nil {
		log.Fatal(err) }

	var results []*User.User

	defer cur.Close(context.TODO())
	
	for cur.Next(context.TODO()) {
		var elem User.User

		if err := cur.Decode(&elem); err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return results
}