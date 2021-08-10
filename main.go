package main

import (
    "context"
    "GO_REST_API/Helper"
	"fmt"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo/options"
    //"encoding/json"
	"log"
)

func main() {

    client := Helper.ConnectDB()
    collection := client.Database("database2").Collection("people")

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

    /*for _, sus := range users{
        fmt.Println(sus)
        fmt.Println()
    }*/

    // Close the cursor once finished
    cur.Close(context.TODO())

    Helper.HandleRequests()
    
    Helper.DisconnectDB(client)
}

/*type Article struct {
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
}*/