package Controller

import (
	"GO_REST_API/SecondLayer"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/gin-gonic/gin"
)

func DelUser(c *gin.Context, collection *mongo.Collection) {
	Service.DelUser_(c, collection)
}