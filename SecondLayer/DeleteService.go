package Service

import (
	"GO_REST_API/ThirdLayer"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/gin-gonic/gin"
    "io/ioutil"
	"log"
	"os"
)

func DelUser_(c *gin.Context, collection *mongo.Collection){

	id := c.Param("id")
	Repository.DelUserRepo(id, c, collection)

	files, err := ioutil.ReadDir("User_data/")
    if err != nil {
        log.Fatal(err)
    }
 
    for _, f := range files {
		if (f.Name() == id) {
			var path string = "User_data/" + f.Name()
			e := os.Remove(path)
			if e != nil {
				log.Fatal(e)
			}
		}
    }
}