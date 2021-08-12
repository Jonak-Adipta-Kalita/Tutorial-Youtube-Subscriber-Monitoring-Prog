package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/Jonak-Adipta-Kalita/JAK-Youtube-Subscriber-Monitoring-Prog/youtube"
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
	item, err := youtube.GetSubscribers()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", item)
	
	setupRoutes()
}
