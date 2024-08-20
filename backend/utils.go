package backend

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func GetAndUnmarshalArtists() ([]Artists, error) {
	artists := []Artists{}
	artistsUrl := "https://groupietrackers.herokuapp.com/api/artists"

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(artistsUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to get artists: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&artists)
		if err != nil {
			return nil, fmt.Errorf("failed to decode artists: %w", err)
		}
	} else {
		return nil, fmt.Errorf("recieved a non-200 response code: %d", resp.StatusCode)
	}

	return artists, nil
}

func GetAndUnmarshalLocations(Id int) (Location, error) {
	locations := Locations{}
	location := Location{}
	locationsURL := "https://groupietrackers.herokuapp.com/api/locations"

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(locationsURL)
	if err != nil {
		return location, fmt.Errorf("failed to get locations: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&locations)
		if err != nil {
			return location, fmt.Errorf("failed to decode locations: %w", err)
		}
	} else {
		return location, fmt.Errorf("received a non-200 response code: %d", resp.StatusCode)
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

func GetAndUnmarshalRelation(Id int) (ArtistDetails, error) {
	relation := Relation{}
	artistDetails := ArtistDetails{}

	relationURL := "https://groupietrackers.herokuapp.com/api/relation"
	resp, err := http.Get(relationURL)
	if err != nil {
		return artistDetails, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&relation)
		if err != nil {
			return artistDetails, err
		}
	} else {
		return artistDetails, fmt.Errorf("failed to get relation data: %s", resp.Status)
	}

	for _, v := range relation.Index {
		if v.ID == Id {
			artistDetails = v
		}
	}

	return artistDetails, nil
}
