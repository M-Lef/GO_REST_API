package main

import (
    "GO_REST_API/Helper"
    //"encoding/json" Marshall
)

func main() {

    client := Helper.ConnectDB("mongodb://localhost:8082")
    collection := client.Database("database2").Collection("people")

	Helper.HandleRequests(collection)
    
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