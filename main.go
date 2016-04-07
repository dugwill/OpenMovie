// Copyright © 2016 Douglas Will
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Issues prints a list of Movie data
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dugwill/MyProjects/omdbapi"
)

//!+
func main() {

	result, err := omdbapi.SearchMovies(os.Args[1:]) // DW: 'result' is a struct of type OmdbSearchResult
	// SearchMovies returns the address of the struct
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s results:\n", result.TotalResults)
	fmt.Printf("\n")

	for _, item := range result.Search { // DW: 'result.Items' is a slice containing the pointers to issues

		fmt.Printf("Title: %s\n", item.Title)
		fmt.Printf("Release Year: %s\n", item.Released)
		fmt.Printf("IMDB ID: %s\n", item.ImdbID)
		fmt.Printf("Type: %s\n", item.Type)
		fmt.Printf("PosterURL: %s\n", item.PosterURL)
		fmt.Printf("\n")

		//poster, err := omdbapi.GetPoster(item.PosterURL, item.Title)
		omdbapi.GetPoster(item.PosterURL, item.Title)
	}
}
