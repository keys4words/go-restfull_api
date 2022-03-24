package handlers

import (
	"booksApi/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateBook(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Creating new book...")
	var book models.Book
	err := json.NewDecoder(request.Body).Decode(&book)
	if err != nil {
		msg := models.Message{Message: "json is invalid"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	var newBookId int = len(models.DB) + 1
	book.ID = newBookId
	models.DB = append(models.DB, book)

	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(book)
}

func GetBookById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("error while parsing", err)
		writer.WriteHeader(400)
		msg := models.Message{Message: "cannot cast this id to int"}
		json.NewEncoder(writer).Encode(msg)
		return
	}
	book, ok := models.FindBookById(id)
	log.Println("Get book with id=", id)
	if !ok {
		writer.WriteHeader(404)
		msg := models.Message{Message: "book with this id doesn't exists"}
		json.NewEncoder(writer).Encode(msg)
	} else {
		writer.WriteHeader(200)
		json.NewEncoder(writer).Encode(book)
	}
}
func UpdateBookById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Updating book...")
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("error while parsing", err)
		writer.WriteHeader(400)
		msg := models.Message{Message: "cannot cast this id to int"}
		json.NewEncoder(writer).Encode(msg)
		return
	}
	oldBook, ok := models.FindBookById(id)
	var newBook models.Book
	if !ok {
		writer.WriteHeader(404)
		log.Println("book not found in db, id:", id)
		msg := models.Message{Message: "book with this id doesn't exists"}
		json.NewEncoder(writer).Encode(msg)
		return
	}
	err = json.NewDecoder(request.Body).Decode(&newBook)
	if err != nil {
		msg := models.Message{Message: "json is invalid"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	for k, v := range models.DB {
		if v.ID == oldBook.ID {
			models.DB[k] = newBook
			models.DB[k].ID = oldBook.ID
		}
	}
	msg := models.Message{Message: "book updated"}
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(msg)
}
func DeleteBookById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Deleting book...")
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("error while parsing", err)
		writer.WriteHeader(400)
		msg := models.Message{Message: "cannot cast this id to int"}
		json.NewEncoder(writer).Encode(msg)
		return
	}
	book, ok := models.FindBookById(id)
	if !ok {
		writer.WriteHeader(404)
		msg := models.Message{Message: "book with this id doesn't exists"}
		json.NewEncoder(writer).Encode(msg)
		return
	}
	for k, v := range models.DB {
		if v.ID == book.ID {
			models.DB = append(models.DB[:k], models.DB[k+1:]...)
			break
		}
	}
	msg := models.Message{Message: "book has been successfully deleted"}
	json.NewEncoder(writer).Encode(msg)
}
