package backend

import (
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"text/template"
)

func artistsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		artists := GetAndUnmarshalArtists()
		tempFile := "frontend/index.html"
		temp, err := template.ParseFiles(tempFile)
		if err != nil {
			log.Fatal(err)
		}
		err = temp.Execute(w, artists)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func artistDetailsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		queryParams := r.URL.Query()
		idValue := queryParams.Get("id")
		ID, err := strconv.Atoi(idValue)
		if err != nil {
			log.Fatal(err)
		}
		location, err := GetAndUnmarshalLocations(ID)
		if err != nil {
			log.Fatal(err)
		}
		tempFile := filepath.Join("frontend", "artist.html")
		temp, err := template.ParseFiles(tempFile)
		if err != nil {
			log.Fatal(err)
		}
		err = temp.Execute(w, location)
		if err != nil {
			log.Fatal(err)
		}
	}
}
