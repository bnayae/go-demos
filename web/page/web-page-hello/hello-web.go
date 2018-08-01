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
	fmt.Println("Serving on ", url)
	log.Fatal(http.ListenAndServe(serveOn, nil))
}

func handleRequests(w http.ResponseWriter, req *http.Request) {
	log.Println("serving", req.URL)
	fmt.Fprint(w, "Hello ", time.Now())
}
