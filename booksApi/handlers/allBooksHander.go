package handlers

import (
	"booksApi/models"
	"encoding/json"
	"log"
	"net/http"
)

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Context-Type", "application/json")
}
func GetAllBooks(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Get info about all books")
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(models.DB)
}
