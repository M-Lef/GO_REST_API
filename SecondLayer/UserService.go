package Service

import (
	"GO_REST_API/ThirdLayer"
	"GO_REST_API/User_Data"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"log"
	"os"
	"fmt"
)

func GetUser_(c *gin.Context, collection *mongo.Collection) User.User{
	id := c.Param("id")
	return Repository.FindUser(id, c, collection)
}

func GetUsers_(c *gin.Context, collection *mongo.Collection) []*User.User{
	return Repository.FindUsers(c, collection)
}

func PutUser_(c *gin.Context, collection *mongo.Collection) {
	id := c.Param("id")
	var user User.User
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = bson.UnmarshalJSON(jsonData, &user)
	if err != nil {
		log.Fatal(err)
	}
	
	database_user := Repository.FindUsers(c, collection)
	for _ , a := range database_user {
		if (a.ID == id) {
			var path string = "User_data/" + id
			content, _ := ioutil.ReadFile(path)
			if (a.Data != string(content)) {
				fmt.Println("i am here")

				err = os.Remove(path)
				if err != nil {
					log.Fatal(err)
				}
				err = ioutil.WriteFile(path, []byte(a.Data), 0755)
				if err != nil {
					log.Fatal(err)
				}
			}
			Repository.PutUserRepo(user, c, collection)
			break
		}
	}	
}