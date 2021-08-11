package Service

import (
	"GO_REST_API/ThirdLayer"
	"GO_REST_API/Users"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/gin-gonic/gin"
	"fmt"
)

func GetUser_(c *gin.Context, collection *mongo.Collection) User.User{
	id := c.Param("id")
	fmt.Println(id)
	return Repository.FindUser(id, c, collection)
}

func GetUsers_(c *gin.Context, collection *mongo.Collection) []*User.User{
	return Repository.FindUsers(c, collection)
}