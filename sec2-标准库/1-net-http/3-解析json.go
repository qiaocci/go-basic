package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	// ...
}

func main() {
	res, err := json.Marshal(movies)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("%s\n", res)

	//res2, err := json.MarshalIndent(movies, "", "  ")
	//if err != nil {
	//	log.Panic(err)
	//}
	//fmt.Printf("%s\n", res2)

	var titles []struct{ Title string }
	err = json.Unmarshal(res, &titles)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(titles)
}
