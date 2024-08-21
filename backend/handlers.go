package backend

import (
	"log"
	"net/http"
	"strconv"
)

func artistsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		artists, err := GetAndUnmarshalArtists()
		if err != nil {
			http.Error(w, "Failed to retrieve artists data", http.StatusInternalServerError)
			log.Printf("Error retrieving artists: %v", err)
			return
		}
		renderTemplate(w, "index.html", artists)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func artistDetailsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		queryParams := r.URL.Query()
		idValue := queryParams.Get("id")
		ID, err := strconv.Atoi(idValue)

		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			log.Printf("Error converting id param to int value: %v", err)
			return
		}

		if ID <= 0 || ID > 52 {
			http.Error(w, "ID out of range", http.StatusBadRequest)
			log.Printf("ID out of range: %d", ID)
			return
		}

		relation, err := GetAndUnmarshalRelation(ID)
		if err != nil {
			http.Error(w, "Failed to retrieve relation data", http.StatusInternalServerError)
			log.Printf("Error retrieving relation data: %v", err)
			return
		}

		renderTemplate(w, "artist.html", relation)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
