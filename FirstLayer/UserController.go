package Controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"GO_REST_API/SecondLayer"
)

func GetUser(c *gin.Context, collection *mongo.Collection) {
	Service.GetUser_(c, collection)
}

func GetUsers(c *gin.Context, collection *mongo.Collection) {
	Service.GetUsers_(c, collection)
}