package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	port string = "8080"
	db   []Pizza
)

type Pizza struct {
	ID      int     `json:"id"`
	Dimeter int     `json:"diameter"`
	Price   float64 `json:"price"`
	Title   string  `json:"title"`
}

func FindPizzaById(id int) (Pizza, bool) {
	var pizza Pizza
	var found bool
	for _, p := range db {
		if p.ID == id {
			pizza = p
			found = true
			break
		}
	}
	return pizza, found
}

func GetAllPizzas(writer http.ResponseWriter, request *http.Request) {

}
func GetPizzaById(writer http.ResponseWriter, request *http.Request) {

}

func main() {
	log.Println("Trying to start REST Pizza API...")
	router := mux.NewRouter()
	router.HandleFunc("/pizzas", GetAllPizzas).Methods("GET")
	router.HandleFunc("/pizza/{id}", GetPizzaById).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
