package main

// https://medium.com/@ScullWM/golang-http-server-for-pro-69034c276355
// http://www.gorillatoolkit.org/pkg/mux
// https://medium.com/@matryer/writing-middleware-in-golang-and-how-go-makes-it-so-much-fun-4375c1246e81

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

const dateFormat = "2006-01-02 15:04:05"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", halloHandler)
	r.HandleFunc("/query/{category}", queryCategoryHandler)
	r.HandleFunc("/query/{category}/{id:[0-9]+}", queryCategoryAndIDHandler)
	// Groups can be used inside patterns, as long as they are non-capturing
	r.HandleFunc("/query-enum/{enum:(?:asc|desc|new)}", queryEnumHandler)
	http.Handle("/", r)

	serveOn := "localhost:7771"
	fmt.Println("Serving on ", serveOn)
	srv := &http.Server{
		Handler: r,
		Addr:    serveOn,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func halloHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("serving", req.URL)
	fmt.Fprint(w, "Hello ", time.Now().Format(dateFormat))
}

func queryCategoryHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	category := vars["category"]
	log.Println("serving", req.URL, " category = ", category)
	fmt.Fprint(w, "Query category = ", category, " at ", time.Now().Format(dateFormat))
}

func queryCategoryAndIDHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	category := vars["category"]
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	log.Println("serving", req.URL, " category = ", category, " id = ", id)
	fmt.Fprint(w, "Query category = ", category, " id = ", id, " at ", time.Now().Format(dateFormat))
}

func queryEnumHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	enum := vars["enum"]
	log.Println("serving", req.URL, " enum = ", enum)
	fmt.Fprint(w, "Query enum = ", enum, " at ", time.Now().Format(dateFormat))
}
