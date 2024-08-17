package main

import (
	"groupie-tracker-complete/backend"
	"net/http"
)

func main() {
	server := http.Server{
		Addr:    "192.168.1.6:8080",
		Handler: backend.RegisterRoutes(),
	}
	server.ListenAndServe()
}
