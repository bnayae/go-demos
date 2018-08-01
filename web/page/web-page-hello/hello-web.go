package main

// https://www.youtube.com/watch?v=5bYO60-qYOI

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/hello", handleRequests)
	serveOn := "localhost:7777"
	url := "http://" + serveOn + "/hello"
	fmt.Println("Serving on ", url, " you should send query in the form of ?q=")
	log.Fatal(http.ListenAndServe(serveOn, nil))
}

func handleRequests(w http.ResponseWriter, req *http.Request) {
	q := req.FormValue("q")
	message := " no query"
	if q != "" {
		message = " query is [" + q + "]"
	} else {
		http.Error(w, "missing 'q' URL parameter", http.StatusBadRequest)
		return
	}
	log.Println("serving", req.URL)
	fmt.Fprint(w, "Hello ", message, " ", time.Now().Format("2006-01-02 15:04:05"))
}
