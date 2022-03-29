package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"quadEquation/handlers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

const (
	API_PREFIX string = "api/v1"
)

var (
	port        string
	grabPrefix  string = API_PREFIX + "/grab"  //api/v1/item/
	solvePrefix string = API_PREFIX + "/solve" //api/v1/items/
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("cannot find .env file:", err)
	}
	port = os.Getenv("app_port")
}

func main() {
	fmt.Println("Starting quatEq API on port: ", port)
	router := mux.NewRouter()
	router.HandleFunc(grabPrefix, handlers.ReadData).Methods("POST")
	router.HandleFunc(solvePrefix, handlers.SolveEquation).Methods("GET")

	fmt.Println("Successfully started...")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
