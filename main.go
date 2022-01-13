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
	// fmt.Fprintf(w, "Detail for course "+params["courseid"])
	// fmt.Fprintf(w, "\n")
	// fmt.Fprintf(w, r.Method)
	if r.Method == "GET" {
		if _, ok := courses[params["courseid"]]; ok {
			json.NewEncoder(w).Encode(courses[params["courseid"]])
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - No course found"))
		}
	}

	if r.Method == "DELETE" {
		if _, ok := courses[params["courseid"]]; ok {
			delete(courses, params["courseid"])
			w.WriteHeader(http.StatusNoContent)
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - No course found"))
		}
	}

	if r.Header.Get("Content-type") == "application/json" {

		if r.Method == "POST" {
			var newCourse courseInfo
			reqBody, err := ioutil.ReadAll(r.Body)
			if err == nil {
				json.Unmarshal(reqBody, &newCourse)
				if newCourse.Title == "" {
					w.WriteHeader(http.StatusUnprocessableEntity)
					w.Write([]byte("422 - Please supply course information in JSON format"))
					return
				}
				if _, ok := courses[params["courseid"]]; !ok {
					courses[params["courseid"]] = newCourse
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("201 - Course added: " + params["courseid"]))
				} else {
					w.WriteHeader(http.StatusConflict)
					w.Write([]byte("409 - Duplicate course ID"))
				}
			} else {
				w.WriteHeader(http.StatusUnprocessableEntity)
				w.Write([]byte("422 - Please supply course information in JSON format"))
			}
		}

		if r.Method == "PUT" {
			var newCourse courseInfo
			reqBody, err := ioutil.ReadAll(r.Body)
			if err == nil {
				json.Unmarshal(reqBody, &newCourse)

				if newCourse.Title == "" {
					w.WriteHeader(http.StatusUnprocessableEntity)
					w.Write([]byte("422 - Please supply course information in JSON format"))
					return
				}

				if _, ok := courses[params["courseid"]]; !ok {
					courses[params["courseid"]] = newCourse
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("201 - Course added: " + params["courseid"]))
				} else {
					courses[params["courseid"]] = newCourse
					w.WriteHeader(http.StatusNoContent)
				}
			} else {
				w.WriteHeader(http.StatusUnprocessableEntity)
				w.Write([]byte("422 - Please supply course information in JSON format"))
			}
		}
	}
}

func main() {
	courses = make(map[string]courseInfo)

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/", home)

	router.HandleFunc("/api/v1/courses", allcources)
	router.HandleFunc("/api/v1/courses/{courseid}", course).Methods("GET", "PUT", "POST", "DELETE")

	fmt.Println("Listening on port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
