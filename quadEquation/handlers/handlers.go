package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"quadEquation/data"
	"quadEquation/utils"
)

func setHeader(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

func ReadData(writer http.ResponseWriter, request *http.Request) {
	setHeader(writer)
	log.Println("ReadData func started...")

	var args data.QuadraticEquation
	err := json.NewDecoder(request.Body).Decode(&args)
	if err != nil {
		msg := utils.Msg{Msg: "provided json file is invaldid"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	if len(data.DB) != 0 {
		data.DB = data.DB[:0]
	}
	data.DB = append(data.DB, &args)
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(data.DB)
}

func SolveEquation(writer http.ResponseWriter, request *http.Request) {
	setHeader(writer)
	log.Println("SolveEquation func started...")

	for _, v := range data.DB {
		if !v.Calculated {
			utils.CalcNumRoots()
		}
	}
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(data.DB)
}
