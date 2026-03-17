package main

import (
	"fmt"
	"io"
	"net/http"
)

type Source string
type Category string

const (
	CatSport  Category = "Sport"
	CatMoto   Category = "Moto"
	CatGossip Category = "Gossip"
	CatInfo   Category = "Info"
)

const (
	Onet   Source = "Onet.pl"
	WP     Source = "Wp.pl"
	O2     Source = "o2.pl"
	Gazeta Source = "gazeta.pl"
)

type Article struct {
	Title       string
	Source      Source
	Category    Category
	Link        string
	Description string
}

func main() {
	url := "https://www.onet.pl/"
	result, err := fetch(url)

	if err != nil {
		fmt.Println("Error during fetch data", err)
		return
	}

	fmt.Println("HTML starts with:")
	fmt.Println(string(result)[:500])
}

func fetch(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", fmt.Errorf("Connection error: %s", resp.Status)
	}

	defer resp.Body.Close()

	fmt.Println("Response code", resp.Status)

	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)

		if err != nil {
			return "", fmt.Errorf("Response read error: %s", resp.Status)
		}

		return string(body), nil
	}

	return "", fmt.Errorf("Fetch error: %s", err)
}
