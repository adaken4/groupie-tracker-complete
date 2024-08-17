package backend

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetAndUnmarshalArtists() []Artists {
	artists := []Artists{}

	artistsUrl := "https://groupietrackers.herokuapp.com/api/artists"
	resp, err := http.Get(artistsUrl)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {

		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&artists)
		if err != nil {
			log.Fatalln(err)
		}
	}
	return artists
}

func GetAndUnmarshalLocations() (Locations, error) {
	locations := Locations{}

	locationsURL := "https://groupietrackers.herokuapp.com/api/locations"
	resp, err := http.Get(locationsURL)
	if err != nil {
		return locations, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&locations)
		if err != nil {
			return locations, err
		}
	} else {
		return locations, fmt.Errorf("failed to get data: %s", resp.Status)
	}

	return locations, nil
}
