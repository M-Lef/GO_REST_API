package Controller

import (
	"GO_REST_API/SecondLayer"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/gin-gonic/gin"
)

func LogUser(c *gin.Context, collection *mongo.Collection) {
	user := Service.LogUser_(c, collection)
    c.JSON(200, user)
}