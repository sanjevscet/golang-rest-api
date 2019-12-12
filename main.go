package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HomePage Endpoint hi Sanjeev!!")
}

func handleRequest() {
	http.HandleFunc("/", handleHomePage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequest()
}
