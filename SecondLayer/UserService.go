package Service

import (
	"GO_REST_API/ThirdLayer"
	"GO_REST_API/User_Data"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/gin-gonic/gin"
)

func GetUser_(c *gin.Context, collection *mongo.Collection) User.User{
	id := c.Param("id")
	return Repository.FindUser(id, c, collection)
}

func GetUsers_(c *gin.Context, collection *mongo.Collection) []*User.User{
	return Repository.FindUsers(c, collection)
}