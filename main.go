package main

import (
	"fmt"
	"net/http"
	"log"
)

func homePage(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintf(w, "Hello World!!")
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func main() {
	fmt.Println("Monitoring Your Youtube Subscribers")
	setupRoutes()
}
