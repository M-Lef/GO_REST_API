package Helper

import (
	"GO_REST_API/FirstLayer"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/gin-gonic/gin"
	"encoding/json"
	"net/http"
	"context"
	"fmt"
	"log"
)

// HandleRequests : Here we are creating our endpoints
func HandleRequests(collection *mongo.Collection, port string) {
    router := gin.Default()

	router.GET("/", homePage)

	router.POST("/add/users", func(c *gin.Context) {Controller.AddUser(c, collection)})
	router.POST("/login", homePage)

	router.DELETE("/delete/user/:id", func(c *gin.Context) {Controller.DelUser(c, collection)})

	router.GET("user/:id", func(c *gin.Context) {Controller.GetUser(c, collection)})
	router.GET("/users/list", func(c *gin.Context) {Controller.GetUsers(c, collection)})

	//Patch method remplace method UPDATE
	router.PATCH("user/:id", homePage)
    
    router.Run(":" + port)
}

func homePage(c *gin.Context){
    jsonData := []byte(`Welcome to homepage`)
    c.Data(http.StatusOK, "application/json", jsonData)
}

// ConnectDB : This is helper function to connect mongoDB
func ConnectDB(uri string) *mongo.Client {

	clientOptions := options.Client().ApplyURI(uri)

    // Connect to MongoDB
    client, err := mongo.Connect(context.TODO(), clientOptions)

    if err != nil {
        log.Fatal(err)
    }

    err = client.Ping(context.TODO(), nil)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connected to MongoDB!")

    return client
}

// DisconnectDB : This is helper function to disconnect mongoDB
func DisconnectDB(client *mongo.Client){
	err := client.Disconnect(context.TODO())

    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Connection to MongoDB closed.")
}

// ErrorResponse : This is error model.
type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"message"`
}

// GetError : This is helper function to prepare error model.
// If you want to export your function. You must to start upper case function name. Otherwise you won't see your function when you import that on other class.
func GetError(err error, w http.ResponseWriter) {

	log.Fatal(err.Error())
	var response = ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	w.Write(message)
}