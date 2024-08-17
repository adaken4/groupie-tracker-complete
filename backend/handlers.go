package backend

import (
	"log"
	"net/http"
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
