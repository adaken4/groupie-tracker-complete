package backend

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var client = &http.Client{
	Timeout: 10 * time.Second,
}

const APIURL = "https://groupietrackers.herokuapp.com/api"

func GetAndUnmarshalArtists() ([]Artists, error) {
	artists := []Artists{}

	jsonData, err := getJSONData(APIURL + "/artists")
	if err != nil {
		return artists, err
	}
	json.Unmarshal(jsonData, &artists)

	return artists, nil
}

func GetAndUnmarshalLocations(Id int) (Location, error) {
	locations := Locations{}
	location := Location{}

	jsonData, err := getJSONData(APIURL + "/locations")
	if err != nil {
		return location, err
	}
	json.Unmarshal(jsonData, &locations)

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

	jsonData, err := getJSONData(APIURL + "/dates")
	if err != nil {
		return date, err
	}
	json.Unmarshal(jsonData, &dates)
	for _, v := range dates.Index {
		if v.ID == Id {
			date = v
		}
	}

	return date, nil
}

func GetAndUnmarshalRelation(Id int) (ArtistDetails, error) {
	relation := Relation{}
	artistDetails := ArtistDetails{}

	jsonData, err := getJSONData(APIURL + "/relation")
	if err != nil {
		return artistDetails, err
	}

	json.Unmarshal(jsonData, &relation)
	for _, v := range relation.Index {
		if v.ID == Id {
			artistDetails = v
		}
	}

	return artistDetails, nil
}

func getJSONData(url string) (json.RawMessage, error) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get json data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received a non-200 response code: %d", resp.StatusCode)
	}

	var body json.RawMessage
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return nil, fmt.Errorf("failed to decode json data: %w", err)
	}

	return body, nil
}
