package main


import (
    //"fmt"
    //"log"
	"net/http"
    "github.com/gin-gonic/gin"
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

/*func allArticles(w http.ResponseWriter, r *http.Request){
	articles := Articles{
		Article{Title:"Title Test", Desc:"Descritpion test", Content:"Salut a tous"},
	}
    json.NewEncoder(w).Encode(articles)
}*/

/*func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}*/

/*func allArticles(){
	articles := Articles{
		Article{Title:"Title Test", Desc:"Descritpion test", Content:"Salut a tous"},
	}
    fmt.Println("Endpoint Hit: AllArticles")
    json.NewEncoder(w).Encode(articles)
}*/

func homePage(c *gin.Context){
    jsonData := []byte(`{"msg":"welcome to homepage"}`)
    c.Data(http.StatusOK, "application/json", jsonData)
}

func handleRequests() {
    router := gin.Default()
    router.GET("/", homePage)

	router.GET("/all", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "articles",
		})
    })

    /*router.GET("/", homePage)
    router.GET("/articles", allArticles)*/
    
    router.Run()
    
   /* myRouter := mux.NewRouter().StrictSlash(true)

    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/all", allArticles)

	log.Fatal(http.ListenAndServe(":8081", myRouter))*/
}

func main() {
    handleRequests()
}