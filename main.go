package main

import (
	"fmt"
	"net/http"
	"log"
	//"github.com/Jonak-Adipta-Kalita/JAK-Youtube-Subscriber-Monitoring-Prog/youtube"
	"github.com/Jonak-Adipta-Kalita/JAK-Youtube-Subscriber-Monitoring-Prog/websocket"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Monitoring Subscriber Counts")
}

func stats(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}
	
	go websocket.Writer(ws)
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/stats", stats)
	
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	fmt.Println("Monitoring Your Youtube Subscribers")
	
	setupRoutes()
}
