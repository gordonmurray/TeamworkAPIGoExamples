package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type tags struct {
	Tags   []struct {
		Name  string `json:"name"`
		ID    string `json:"id"`
		Color string `json:"color"`
	} `json:"tags"`
	STATUS string `json:"STATUS"`
}

func main() {

	url := fmt.Sprintf("https://xxxxxx.teamwork.com/tags.json")

	req, err := http.NewRequest("GET", url, nil)

	req.SetBasicAuth("xxxxxx", "")

	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	// HTTP client
	client := &http.Client{}

	// Send the request via a client
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	// close resp.Body
	defer resp.Body.Close()

	// Fill the record with the data from the JSON
	var record tags

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	if (record.STATUS == "OK") {
		for _, element := range record.Tags {
			fmt.Println("Tag " + element.ID + " is " + element.Name + " and its colour is " + element.Color)
		}
	}
}
