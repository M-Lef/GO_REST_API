package Service

import (
	"GO_REST_API/ThirdLayer"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/gin-gonic/gin"
	"GO_REST_API/User_Data"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"fmt"
)

func AddUser_(c *gin.Context, collection *mongo.Collection){

	var users[] User.User
	file, _, err := c.Request.FormFile("file")
    if err != nil {
        return
    }
    defer file.Close()
	buf := make([]byte, 1024)
	var str string
    for {
        len, _ := file.Read(buf)
        if len == 0 {
            break
        }
		str += string(buf)
	}
	err = bson.UnmarshalJSON([]byte(str),&users)
	if err != nil {
		panic(err)
	}
	var new_users[] User.User

	for _ , a := range users {
		database_user := Repository.FindUsers(c, collection)
		var pres bool = true
		for _ , b := range database_user {
			if (a.ID == b.ID) {
				pres = false
				break
			}
		}
		if (pres) {
			path := "User_data/" + a.ID
			err = ioutil.WriteFile(path, []byte(a.Data), 0755)
			if err != nil {
				fmt.Printf("Unable to write file: %v", err)
			}
			new_users = append(new_users, a)
		}
	}

	Repository.PostUser(new_users, c, collection)
}