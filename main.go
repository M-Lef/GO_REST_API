package main

import (
    "GO_REST_API/Helper"
)

func main() {

    // You should change the URI for connecting to your MongoDB server
    client := Helper.ConnectDB("mongodb://localhost:8082")

    // You should change the name of the database and the collection
    collection := client.Database("database2").Collection("people")

    // You can choose on what port the request are going to be made
	Helper.HandleRequests(collection, "8080")
    
    Helper.DisconnectDB(client)
}
