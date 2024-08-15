package backend

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func PrintArtist(ID int) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var artists []Artists
	err = json.Unmarshal(respBody, &artists)
	for _, v := range artists {
		if ID == v.ID {
			fmt.Println(v)
		}
	}
}
