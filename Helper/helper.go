package Helper

import (
	"context"
	"GO_REST_API/FirstLayer"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/gin-gonic/gin"
)

// HandleRequests : Here we are creating our endpoints
func HandleRequests(collection *mongo.Collection) {
    router := gin.Default()

	router.GET("/", homePage)

	router.POST("/add/user", homePage)
	router.POST("/login", homePage)

	router.DELETE("/delete/user/:id", homePage)

	router.GET("user/:id", func(c *gin.Context) {Controller.GetUser(c, collection)})
	router.GET("/users/list", func(c *gin.Context) {Controller.GetUsers(c, collection)})

	//Patch method remplace method UPDATE
	router.PATCH("user/:id", homePage)
    
    router.Run(":8080")
}

func homePage(c *gin.Context){
    jsonData := []byte(`Welcome to homepage`)
    c.Data(http.StatusOK, "application/json", jsonData)
}

// ConnectDB : This is helper function to connect mongoDB
func ConnectDB() *mongo.Client {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:8082")

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

/*func SuS {

	findOptions := options.Find()
	findOptions.SetLimit(20)

	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	//cur, err := collection.Find(context.Background(), bson.D{},)

	if err != nil {
		log.Fatal(err) }

	var users Helper.Users
	var user Helper.User

	for cur.Next(context.TODO()) {
		
		elem := &bson.D{}

		if err := cur.Decode(elem); err != nil {
				log.Fatal(err)
		}

		fmt.Println(elem)
		fmt.Println()

		users = append(users, user)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	for _, sus := range users{
		fmt.Println(sus)
		fmt.Println()
	}

	// Close the cursor once finished
	cur.Close(context.TODO())
}*/