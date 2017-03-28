package main

import (
	"changelog"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	r := mux.NewRouter()
	r.HandleFunc("/", changelog.Handler)

	var handler http.Handler
	handler = handlers.CORS()(r)
	handler = handlers.CompressHandler(handler)

	// Bind to a port and pass our router in
	http.ListenAndServe(":"+port, handler)
}
