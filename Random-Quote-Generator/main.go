package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Quote struct {
	Text   string `json:"content"`
	Author string `json:"author"`
}

func fetchRandomQuote() (Quote, error) {
	apiUrl := "https://api.quotable.io/random"

	response, err := http.Get(apiUrl)
	if err != nil {
		return Quote{}, err
	}
	defer response.Body.Close()

	fmt.Println("API Response Status:", response.Status)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return Quote{}, err
	}
	fmt.Println("Response Body:", string(body))

	var quote Quote
	if err := json.Unmarshal(body, &quote); err != nil {
		return Quote{}, err
	}

	fmt.Println("Fetched Quote:", quote.Text)

	return quote, nil
}

func main() {
	rand.Seed(time.Now().UnixNano())

	quote, err := fetchRandomQuote()
	if err != nil {
		log.Fatalf("Error fetching quote: %v", err)
	}

	fmt.Println("Random Quote:")
	fmt.Printf("'%s' - %s\n", quote.Text, quote.Author)
}
