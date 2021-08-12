package Controller

import (
	"GO_REST_API/SecondLayer"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context, collection *mongo.Collection) {
	Service.AddUser_(c, collection)
}