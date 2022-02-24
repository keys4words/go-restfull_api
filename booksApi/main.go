package main

import (
	"log"
	"net/http"
	"os"

	"booksApi/utils"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

const (
	apiPrefix string = "/api/v1"
)

var (
	port                  string
	bookResourcePrefix    string = apiPrefix + "/book"
	allBookResourcePrefix string = apiPrefix + "/books"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("could not find .env file", err)
	}
	port = os.Getenv("app_port")
}

func main() {
	log.Println("Books API started at port", port)
	router := mux.NewRouter()

	utils.BuildBookResource(router, bookResourcePrefix)
	utils.BuildAllBooksResource(router, allBookResourcePrefix)

	log.Println("Router initializing successfully!")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
