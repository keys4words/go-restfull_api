package utils

import (
	"booksApi/handlers"
	"github.com/gorilla/mux"
)

func BuildAllBooksResource(router *mux.Router, prefix string) {
	router.HandleFunc(prefix+"/{id}", handlers.GetBookById).Methods("GET")
	router.HandleFunc(prefix+"/{id}", handlers.UpdateBookById).Methods("PUT")
	router.HandleFunc(prefix+"/{id}", handlers.DeleteBookById).Methods("DELETE")
	router.HandleFunc(prefix, handlers.CreateBook).Methods("POST")
}
func BuildBookResource(router *mux.Router, prefix string) {
	router.HandleFunc(prefix, handlers.GetAllBooks).Methods("GET")
}
