package Controller

import (
	"GO_REST_API/SecondLayer"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context, collection *mongo.Collection) {
	user := Service.GetUser_(c, collection)
    c.JSON(200, user)
}

func GetUsers(c *gin.Context, collection *mongo.Collection) {
	users := Service.GetUsers_(c, collection)
	c.JSON(200, users)
}

func PutUser(c *gin.Context, collection *mongo.Collection) {
	Service.PutUser_(c, collection)
}