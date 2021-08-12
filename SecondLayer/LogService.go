package Service

import (
	"GO_REST_API/ThirdLayer"
	"GO_REST_API/User_Data"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"strings"
	"log"
)

func LogUser_(c *gin.Context, collection *mongo.Collection) User.User{

	var user User.User

	id := c.Param("id")
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatal(err)
	}

	s := strings.Split(string(jsonData), "\n")

	id, password := s[0], s[1]
	
	database_user := Repository.FindUsers(c, collection)
	for _ , a := range database_user {
		if (a.ID == id) {
			err = bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(password))
    		if err == nil {
				user = *a
				break
			}
		}
	}
	return user
}