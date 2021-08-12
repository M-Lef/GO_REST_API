package main

import (
    "GO_REST_API/Helper"
)

func main() {

    client := Helper.ConnectDB("mongodb://localhost:8082")
    collection := client.Database("database2").Collection("people")

	Helper.HandleRequests(collection)
    
    Helper.DisconnectDB(client)
}