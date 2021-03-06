package main

// Response Snippets
// https://www.alexedwards.net/blog/golang-response-snippets

// Http Client
// https://golang.org/pkg/net/http/

// Routing
// https://medium.com/@ScullWM/golang-http-server-for-pro-69034c276355
// http://www.gorillatoolkit.org/pkg/mux
// https://medium.com/@matryer/writing-middleware-in-golang-and-how-go-makes-it-so-much-fun-4375c1246e81

// SVG
// https://github.com/ajstarks/svgo

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"time"

	svg "github.com/ajstarks/svgo"
	"github.com/gorilla/mux"
)

const dateFormat = "2006-01-02 15:04:05"

const imageURL = "https://source.unsplash.com/1600x900?dog"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/header-only", headerOnlyHandler)
	r.HandleFunc("/plan-text", plainTextHandler)
	r.HandleFunc("/json", jsonHandler)
	r.HandleFunc("/json-indent", jsonHandler)
	r.HandleFunc("/xml", xmlHandler)
	r.HandleFunc("/xml-indent", xmlHandler)
	r.HandleFunc("/file", fileHandler)
	r.HandleFunc("/file-fix", fileFixHandler)
	r.HandleFunc("/svg", svgHandler)
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
	fmt.Println("8." + url + "file-fix")
	fmt.Println("9." + url + "svg")

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

	response, err := http.Get(imageURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer response.Body.Close()
	logRequest(req) // TODO: Middleware

	body, err := ioutil.ReadAll(response.Body) // ReadCloser -> []byte
	if err != nil {
		panic(err.Error())
	}

	reader := bytes.NewReader(body) // ReadSeeker
	http.ServeContent(w, req, "image", time.Now(), reader)
}

func fileFixHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("loading image ...")

	logRequest(req) // TODO: Middleware
	fp := path.Join("images", "squirrel.jpg")
	http.ServeFile(w, req, fp)
}

func svgHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("creating SVG ...")

	logRequest(req) // TODO: Middleware
	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	s.Start(500, 500)
	s.Circle(100, 100, 100, "fill:none;stroke:black")
	s.Circle(200, 200, 100, "fill:none;stroke:black")
	s.Circle(300, 300, 100, "fill:none;stroke:black")
	s.Circle(400, 400, 100, "fill:none;stroke:black")
	s.Grid(0, 0, 500, 500, 50, "fill:none;stroke:black")
	s.End()
}

func logRequest(req *http.Request) {
	log.Println("serving", req.URL, " ", req.Method, " "+time.Now().Format(dateFormat)) // TODO: Middleware
}
