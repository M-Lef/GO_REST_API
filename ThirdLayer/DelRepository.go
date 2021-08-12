package Repository

import (
	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/gin-gonic/gin"
	//"context"
)

func DelUserRepo(id string, c *gin.Context, collection *mongo.Collection){
	//f := bson.M{"id": bson.M{"$eq": id}}
	//collection.DeleteMany(context.Background(), f)
}