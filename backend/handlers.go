package backend

import (
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"text/template"
)

// type ArtistDetails struct {
// 	Location       Location
// 	Date           Date
// 	DatesLocations DatesLocations
// }

func artistsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		artists, err := GetAndUnmarshalArtists()
		if err != nil {
			http.Error(w, "Failed to retrieve artists data", http.StatusInternalServerError)
			log.Printf("Error retrieving artists: %v", err)
			return
		}
		tempFile := filepath.Join("frontend", "index.html")
		temp, err := template.ParseFiles(tempFile)
		if err != nil {
			http.Error(w, "Failed to load template: %v", http.StatusInternalServerError)
			log.Printf("Error parsing template: %v", err)
			return
		}
		err = temp.Execute(w, artists)
		if err != nil {
			http.Error(w, "Failed to execute template", http.StatusInternalServerError)
			log.Printf("Error parsing template: %v", err)
		}
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

		// location, err := GetAndUnmarshalLocations(ID)
		// if err != nil {
		// 	http.Error(w, "Failed to retrieve locations data", http.StatusInternalServerError)
		// 	log.Printf("Error retrieving artist locations: %v", err)
		// 	return
		// }

		// date, err := GetAndUnmarshalDates(ID)
		// if err != nil {
		// 	http.Error(w, "Failed to retrieve dates data", http.StatusInternalServerError)
		// 	log.Printf("Error retrieving artist dates: %v", err)
		// 	return
		// }

		relation, err := GetAndUnmarshalRelation(ID)
		if err != nil {
			http.Error(w, "Failed to retrieve relation data", http.StatusInternalServerError)
			log.Printf("Error retrieving relation data: %v", err)
			return
		}

		tempFile := filepath.Join("frontend", "artist.html")
		temp, err := template.ParseFiles(tempFile)
		if err != nil {
			log.Fatal(err)
		}

		// Pass the combined data to the template
		err = temp.Execute(w, relation)
		if err != nil {
			log.Fatal(err)
		}
	}
}
