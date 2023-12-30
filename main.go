package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"math/rand"
	"strings"
)

func main() {
    // Parse command-line arguments
    filePath := flag.String("file", "", "Path to the code file to analyze")
    flag.Parse()

    // Read code file and count lines
    file, err := os.ReadFile(*filePath)
    if err != nil {
        panic(err) // Replace with proper error handling
    }
    lines := strings.Split(string(file), "\n")
    numLines := len(lines)

    // Grant experience and trigger encounters
    xpPerLine := 10
    xpGained := numLines * xpPerLine
    fmt.Printf("You gained %d XP!\n", xpGained)

    encounterChance := 10
    if numLines%encounterChance == 0 {
        encounterPokemon(xpGained)
    }

    // (Coming up) Fetch Pokémon data from PokeAPI and display output
}

func fetchRandomPokemon(xp int) struct {
	Name string `json: "name"`
	Url string `json: "url"`}{
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon?offset=%d&limit=10", xp%100)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var pokemonResults struct {
		Results []struct {
			Name string `json: "name"`
			Url string `json: "url"`
		} `json: "results"`
	}
	err = json.NewDecoder(res.Body).Decode(&pokemonResults)
	if err != nil {
		panic(err)
	}

	randomIndex := rand.Intn(len(pokemonResults.Results))
	randomPokemon := pokemonResults.Results[randomIndex]

	return randomPokemon
}

func encounterPokemon(xp int) {
	randomPokemon := fetchRandomPokemon(xp)

	fmt.Println("A wild Pokemon appears")
	fmt.Printf("It's a %s!\n", randomPokemon.Name)
}