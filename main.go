package main

import (
	"net/http"

	"groupie-tracker-complete/backend"
)

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler: backend.RegisterRoutes(),
	}
	server.ListenAndServe()
}
