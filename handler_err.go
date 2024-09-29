package main

import "net/http"

func HandlerErr(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 404, "Not found")
}
