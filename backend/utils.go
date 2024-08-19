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

func GetAndUnmarshalLocations(Id int) (Location, error) {
	locations := Locations{}
	location := Location{}

	locationsURL := "https://groupietrackers.herokuapp.com/api/locations"
	resp, err := http.Get(locationsURL)
	if err != nil {
		return location, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&locations)
		if err != nil {
			return location, err
		}
	} else {
		return location, fmt.Errorf("failed to get data: %s", resp.Status)
	}
	for _, v := range locations.Index {
		if v.ID == Id {
			location = v
		}
	}

	return location, nil
}

func GetAndUnmarshalDates(Id int) (Date, error) {
	dates := Dates{}
	date := Date{}

	datesURL := "https://groupietrackers.herokuapp.com/api/dates"
	resp, err := http.Get(datesURL)
	if err != nil {
		return date, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&dates)
		if err != nil {
			return date, err
		}
	} else {
		return date, fmt.Errorf("failed to get data: %s", resp.Status)
	}
	for _, v := range dates.Index {
		if v.ID == Id {
			date = v
		}
	}

	return date, nil
}

func GetAndUnmarshalRelation(Id int) (Location, error) {
	locations := Locations{}
	location := Location{}

	locationsURL := "https://groupietrackers.herokuapp.com/api/relation"
	resp, err := http.Get(locationsURL)
	if err != nil {
		return location, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&locations)
		if err != nil {
			return location, err
		}
	} else {
		return location, fmt.Errorf("failed to get data: %s", resp.Status)
	}

	return locations.Index[Id-1], nil
}
