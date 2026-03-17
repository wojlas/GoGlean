package main

import (
	"errors"
	"fmt"
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
	articles := []Article{
		{
			Title:       "Test Onet",
			Source:      "Onet.pla",
			Category:    CatSport,
			Link:        "www.onet.pl",
			Description: "test news onet",
		},
		{
			Title:       "Test Wp",
			Source:      "Wp.pl",
			Category:    CatInfo,
			Link:        "www.wp.pl",
			Description: "test news wp",
		},
		{
			Title:       "Test o2",
			Source:      "o2.pl",
			Category:    CatGossip,
			Link:        "www.02.pl",
			Description: "test news o2",
		},
		{
			Title:       "Test gazeta",
			Source:      "gazeta.pl",
			Category:    CatMoto,
			Link:        "www.gazeta.pl",
			Description: "test news gazeta",
		},
	}

	found, err := findArticleByTitle(articles, "Test Onet")

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Found:", found.Title)
	}
}

func findArticleByTitle(articles []Article, title string) (Article, error) {
	for _, a := range articles {
		if a.Title == title {
			return a, nil
		}
	}

	return Article{}, errors.New("no article found")
}
