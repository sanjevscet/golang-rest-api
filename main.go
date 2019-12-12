package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Article single
type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// Articles list of Article
type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {

	articles := Articles{
		Article{
			Title:   "Test Title",
			Desc:    "Test Description",
			Content: "Hello World",
		},

		Article{
			Title:   "Test Title2",
			Desc:    "Test Description2",
			Content: "Hello Sanjeev",
		},
	}

	fmt.Println("Endpoint Hit:: All Articles")
	json.NewEncoder(w).Encode(articles)
}

func handleHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit::Landing Page")

	fmt.Fprintf(w, "HomePage Endpoint hi Sanjeev!!")
}

func handleRequest() {
	http.HandleFunc("/", handleHomePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequest()
}
