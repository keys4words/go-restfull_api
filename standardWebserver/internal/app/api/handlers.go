package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gorilla/mux"
	"github.com/keys4words/go-restfull_api/standardWebserver/internal/app/middleware"
	"github.com/keys4words/go-restfull_api/standardWebserver/internal/app/models"
)

type Message struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	IsError    bool   `json:"is_error"`
}

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

func (api *API) GetAllArticles(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Get all articles GET /api/v1/articles")
	articles, err := api.storage.Article().SelectAll()
	if err != nil {
		api.logger.Info("Error while Articles.SelectAll - ", err)
		msg := Message{
			StatusCode: 501,
			Message:    "Cannot access to DB",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(articles)
}

func (api *API) GetArticleById(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Get article by id")
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		api.logger.Info("Cannot convert id to int", err)
		msg := Message{
			StatusCode: 400,
			Message:    "Unsupported type of id",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	article, ok, err := api.storage.Article().FindArticleById(id)
	if err != nil {
		api.logger.Info("Error in accessing to db with article id", err)
		msg := Message{
			StatusCode: 500,
			Message:    "Error to access DB with article id",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	if !ok {
		api.logger.Info("Cannot find article with this id in DB")
		msg := Message{
			StatusCode: 404,
			Message:    "Cannot find article with this id in DB",
			IsError:    true,
		}
		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(article)
}
func (api *API) DeleteArticleById(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Delete article by id")
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		api.logger.Info("Cannot convert id to int", err)
		msg := Message{
			StatusCode: 400,
			Message:    "Unsupported type of id",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	_, ok, err := api.storage.Article().FindArticleById(id)
	if err != nil {
		api.logger.Info("Error in accessing to db with article id", err)
		msg := Message{
			StatusCode: 500,
			Message:    "Error to access DB with article id",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	if !ok {
		api.logger.Info("Cannot find article with this id in DB")
		msg := Message{
			StatusCode: 404,
			Message:    "Cannot find article with this id in DB",
			IsError:    true,
		}
		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	_, err = api.storage.Article().DeleteById(id)
	if err != nil {
		api.logger.Info("Troubles with deleting article", err)
		msg := Message{
			StatusCode: 501,
			Message:    "Troubles with deleting article",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(202)
	msg := Message{
		StatusCode: 202,
		Message:    fmt.Sprintf("Article with id %d successfully deleted!", id),
		IsError:    false,
	}
	json.NewEncoder(writer).Encode(msg)
}

func (api *API) PostArticle(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("POST Article /api/v1/articles")
	var article models.Article
	err := json.NewDecoder(req.Body).Decode(&article)
	if err != nil {
		api.logger.Info("Invalid json from client!")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	a, err := api.storage.Article().Create(&article)
	if err != nil {
		api.logger.Info("Troubles while creating new article ", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have troubles in creating new article",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(a)
}
func (api *API) PostUserRegister(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Post User Register")
	var user models.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		api.logger.Info("Invalid json from client!")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	_, ok, err := api.storage.User().FindByLogin(user.Login)
	if err != nil {
		api.logger.Info("Error in accessing to db with user id", err)
		msg := Message{
			StatusCode: 500,
			Message:    "Error to access DB with user id",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	if ok {
		api.logger.Info("User with this id already exists!")
		msg := Message{
			StatusCode: 400,
			Message:    "User with this id already exists",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	userAdded, err := api.storage.User().Create(&user)
	if err != nil {
		api.logger.Info("Error in accessing to db with user id", err)
		msg := Message{
			StatusCode: 500,
			Message:    "Error to access DB with user id",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	msg := Message{
		StatusCode: 201,
		Message:    fmt.Sprintf("User {login:%s} successfully registered!", userAdded.Login),
		IsError:    false,
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(msg)
}

func (api *API) PostToAuth(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Post to auth api")
	var user models.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		api.logger.Info("Invalid json from client!")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	userInDB, ok, err := api.storage.User().FindByLogin(user.Login)
	if err != nil {
		api.logger.Info("Error in accessing to db with user id", err)
		msg := Message{
			StatusCode: 500,
			Message:    "Error to access DB with user id",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	if !ok {
		api.logger.Info("User with this login NOT exists in DB!")
		msg := Message{
			StatusCode: 400,
			Message:    "User with this login NOT exists in DB",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	if userInDB.Password != user.Password {
		api.logger.Info("Invalid credentials to auth")
		msg := Message{
			StatusCode: 404,
			Message:    "Your password is invalid",
			IsError:    true,
		}
		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()
	claims["admin"] = true
	claims["name"] = userInDB.Login
	tokenString, err := token.SignedString(middleware.SecretKey)
	if err != nil {
		api.logger.Info("Cannot claims jwt-token")
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	msg := Message{
		StatusCode: 201,
		Message:    tokenString,
		IsError:    false,
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(msg)
}
