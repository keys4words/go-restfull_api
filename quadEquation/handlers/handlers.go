package handlers

import (
	"encoding/json"
	"net/http"
	"quadEquation/data"
)

func setHeader(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}
func ReadData(writer http.ResponseWriter, request *http.Request) {
	setHeader(writer)

	var args data.QuadraticEquation
	err := json.NewDecoder(request.Body).Decode(&args)
	if err != nil {
		msg := util
	}
}
