package main

// Response Snippets
// https://www.alexedwards.net/blog/golang-response-snippets

// Http Client
// https://golang.org/pkg/net/http/

// Routing
// https://medium.com/@ScullWM/golang-http-server-for-pro-69034c276355
// http://www.gorillatoolkit.org/pkg/mux
// https://medium.com/@matryer/writing-middleware-in-golang-and-how-go-makes-it-so-much-fun-4375c1246e81

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"path"
	"time"

	"github.com/gorilla/mux"
)

const dateFormat = "2006-01-02 15:04:05"

//const imageUrl = "https://source.unsplash.com/1600x900?dog"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/header-only", headerOnlyHandler)
	r.HandleFunc("/plan-text", plainTextHandler)
	r.HandleFunc("/json", jsonHandler)
	r.HandleFunc("/json-indent", jsonHandler)
	r.HandleFunc("/xml", xmlHandler)
	r.HandleFunc("/xml-indent", xmlHandler)
	r.HandleFunc("/file", fileHandler)
	http.Handle("/", r)

	serveOn := "localhost:7771"
	url := "http://" + serveOn + "/"
	fmt.Println("Serving on ", serveOn)

	fmt.Println("Choose one of the following routing options")
	fmt.Println("1." + url + "header-only")
	fmt.Println("2." + url + "plan-text")
	fmt.Println("3." + url + "json")
	fmt.Println("4." + url + "json-indent")
	fmt.Println("5." + url + "xml")
	fmt.Println("6." + url + "xml-indent")
	fmt.Println("7." + url + "file")

	srv := &http.Server{
		Handler: r,
		Addr:    serveOn,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func headerOnlyHandler(w http.ResponseWriter, req *http.Request) {
	logRequest(req) // TODO: Middleware
	w.Header().Set("Server", "A Go Web Server")
	w.WriteHeader(200)
}

func plainTextHandler(w http.ResponseWriter, req *http.Request) {
	logRequest(req) // TODO: Middleware
	w.Write([]byte("OK"))
}

type profile struct {
	ID      int
	Name    string
	Hobbies []string
}

func jsonHandler(w http.ResponseWriter, req *http.Request) {
	logRequest(req) // TODO: Middleware
	profile := profile{42, "Alex", []string{"snowboarding", "programming"}}

	var js []byte
	var err error
	if req.URL.Path == "/json" {
		js, err = json.Marshal(profile)
	} else if req.URL.Path == "/json-indent" {
		js, err = json.MarshalIndent(profile, "", "    ")
	} else {
		http.Error(w, "in correct url option", http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func xmlHandler(w http.ResponseWriter, req *http.Request) {
	logRequest(req) // TODO: Middleware
	profile := profile{42, "Alex", []string{"snowboarding", "programming"}}

	var js []byte
	var err error
	if req.URL.Path == "/xml" {
		js, err = xml.Marshal(profile)
	} else if req.URL.Path == "/xml-indent" {
		js, err = xml.MarshalIndent(profile, "", "    ")
	} else {
		http.Error(w, "in correct url option", http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func fileHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("loading image ...")

	// img, _ := os.Create("image.jpg")
	// defer img.Close()

	// response, err := http.Get(imageUrl)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// defer response.Body.Close()

	// size, _ := io.Copy(img, response.Body)
	// fmt.Println("File size: ", size)

	// logRequest(req) // TODO: Middleware

	// http.ServeFile(w, req, img)

	logRequest(req) // TODO: Middleware
	fp := path.Join("images", "squirrel.jpg")
	http.ServeFile(w, req, fp)
}

func logRequest(req *http.Request) {
	log.Println("serving", req.URL, " ", req.Method, " "+time.Now().Format(dateFormat)) // TODO: Middleware
}
