package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to REST API!")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/", home)

	fmt.Println("Listening on port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}