package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(writer http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5XX error:", msg)
	}

	type errorResponse struct {
		Error string `json:"error"`
	}

	respondWithJson(writer, code, errorResponse{
		Error: msg,
	})
}

func respondWithJson(writer http.ResponseWriter, code int, payload interface{}) {
	data, error := json.Marshal(payload)
	if error != nil {
		log.Printf("failed to marshal json response: payload: %v, error: %v", payload, error)
		writer.WriteHeader(500)
		return
	}
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(code)
	writer.Write(data)
}
