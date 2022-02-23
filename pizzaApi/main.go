package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var (
	port string = "8080"
	db   []Pizza
)

func init() {
	pizza1 := Pizza{
		ID:      1,
		Dimeter: 22,
		Price:   500.54,
		Title:   "Pepperoni",
	}
	pizza2 := Pizza{
		ID:      2,
		Dimeter: 22,
		Price:   680.54,
		Title:   "BBQ",
	}
	pizza3 := Pizza{
		ID:      3,
		Dimeter: 22,
		Price:   450.04,
		Title:   "4Cheese",
	}
	db = append(db, pizza1, pizza2, pizza3)
}

type Pizza struct {
	ID      int     `json:"id"`
	Dimeter int     `json:"diameter"`
	Price   float64 `json:"price"`
	Title   string  `json:"title"`
}

type ErrorMessage struct {
	Message string `json:"message"`
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
	writer.Header().Set("Content-Type", "application/json")
	log.Println("Show all pizzas")
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(db)

}
func GetPizzaById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("client try use invalid id params", err)
		msg := ErrorMessage{Message: "id must be int type"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	log.Println("Trying to send to client pizza with id=", id)
	pizza, ok := FindPizzaById(id)
	if ok {
		writer.WriteHeader(200)
		json.NewEncoder(writer).Encode(pizza)
	} else {
		writer.WriteHeader(404)
		msg := ErrorMessage{Message: "pizza with this id doesn't exist"}
		json.NewEncoder(writer).Encode(msg)
	}
}

func main() {
	log.Println("Trying to start REST Pizza API...")
	router := mux.NewRouter()
	router.HandleFunc("/pizzas", GetAllPizzas).Methods("GET")
	router.HandleFunc("/pizza/{id}", GetPizzaById).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
