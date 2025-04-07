package main

import "net/http"

func handlerReadiness(writer http.ResponseWriter, r *http.Request) {
	respondWithJson(writer, 200, struct{}{})
}
