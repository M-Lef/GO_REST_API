package main

import (
    "fmt"
    "log"
	"net/http"
	"encoding/json"
)

type Article struct {
    Title string `json:"Title"`
    Desc string `json:"desc"`
    Content string `json:"content"`
}

// let's declare a global Articles array
// that we can then populate in our main function
// to simulate a database
type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request){
	articles := Articles{
		Article{Title:"Title Test", Desc:"Descritpion test", Content:"Salut a tous"},
	}
    fmt.Println("Endpoint Hit: AllArticles")
    json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
    log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
    handleRequests()
}