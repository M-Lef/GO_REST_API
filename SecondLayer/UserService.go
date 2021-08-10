package Service

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"fmt"
	"GO_REST_API/ThirdLayer"
)

func GetUser_(c *gin.Context, collection *mongo.Collection){
	id := c.Param("id")
	fmt.Println(id)
	Repository.FindUser(id, c, collection)
}

func GetUsers_(c *gin.Context, collection *mongo.Collection) {
	Repository.FindUsers(c, collection)
}