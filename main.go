package main

import (
    "GO_REST_API/Helper"
)

func main() {

    client := Helper.ConnectDB("mongodb://localhost:8082")
    collection := client.Database("database2").Collection("people")

	Helper.HandleRequests(collection, "8080")
    
    Helper.DisconnectDB(client)
}
