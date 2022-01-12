package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type courseInfo struct {
	Title string `json: "Title"`
}

var courses map[string]courseInfo

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to REST API!")
}

func allcources(w http.ResponseWriter, r *http.Request) {

	kv := r.URL.Query()
	for k, v := range kv {
		fmt.Println(k, v)
	}
	json.NewEncoder(w).Encode(courses)
}

func course(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Fprintf(w, "Detail for course "+params["courseid"])
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, r.Method)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/", home)

	router.HandleFunc("/api/v1/courses", allcources)
	router.HandleFunc("/api/v1/courses/{courseid}", course).Methods("GET", "PUT", "POST", "DELETE")

	fmt.Println("Listening on port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
