package main


import (
	"context"
	"fmt"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo/options"
    //"encoding/json"
	"log"
)
type User []struct {
	ID         string   `json:"id"`
	Password   string   `json:"password"`
	IsActive   bool     `json:"isActive"`
	Balance    string   `json:"balance"`
	Age        string   `json:"age"`
	Name       string   `json:"name"`
	Gender     string   `json:"gender"`
	Company    string   `json:"company"`
	Email      string   `json:"email"`
	Phone      string   `json:"phone"`
	Address    string   `json:"address"`
	About      string   `json:"about"`
	Registered string   `json:"registered"`
	Latitude   float64  `json:"latitude"`
	Longitude  float64  `json:"longitude"`
	Tags       []string `json:"tags"`
	Friends    []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"friends"`
	Data string `json:"data"`
}

type Users []User

func main() {
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

    collection := client.Database("database2").Collection("people")

    findOptions := options.Find()
    findOptions.SetLimit(20)

    cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
    //cur, err := collection.Find(context.Background(), bson.D{},)
    
    if err != nil {
        log.Fatal(err) }

    var users Users
    var user User

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

    /*for _, sus := range users{
        fmt.Println(sus)
        fmt.Println()
    }*/

    // Close the cursor once finished
    cur.Close(context.TODO())
    
    err = client.Disconnect(context.TODO())

    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Connection to MongoDB closed.")

    }



/*package main


import (
	"net/http"
    "github.com/gin-gonic/gin"
)

type Article struct {
    Title string `json:"Title"`
    Desc string `json:"desc"`
    Content string `json:"content"`
}

type Articles []Article

func allArticles(c *gin.Context){
	articles := Articles{
		Article{Title:"Title Test", Desc:"Descritpion test", Content:"Salut a tous"},
    }
    c.JSON(200, articles)
}

func homePage(c *gin.Context){
    jsonData := []byte(`{"msg":"welcome to homepage"}`)
    c.Data(http.StatusOK, "application/json", jsonData)
}

func handleRequests() {
    router := gin.Default()

    router.GET("/", homePage)
	router.GET("/all", allArticles)
    
    router.Run()
}

func main() {
    handleRequests()
}*/