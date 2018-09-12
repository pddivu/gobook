package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {

	type Movie struct {
		Title  string
		Year   int  `json:"released"`
		Color  bool `json:"color,omitempty"`
		Actors []string
	}
	var Movies = []Movie{
		{Title: "Ferdinand2", Year: 2017, Color: true,
			Actors: []string{"Bull", "goat"}},
		{Title: "Ferdinand", Year: 2017,
			Actors: []string{"Bull", "goat"}}}

	fmt.Println(Movies)
	//data, err := json.Marshal(Movies)
	data, err := json.MarshalIndent(Movies, "", "  ")
	if err != nil {
		log.Fatalf("JSON marshalling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	var title []struct{ Title string }
	if err := json.Unmarshal(data, &title); err != nil {
		log.Fatalf("Unmarshalling error: %s", err)

	}
	fmt.Println(title)
}
