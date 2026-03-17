package main

import (
	"fmt"
)

type Article struct {
	Title       string
	Source      string
	Category    string
	Link        string
	Description string
}

func main() {
	firstNews := Article{
		Title:       "Test Onet",
		Source:      "Onet.pl",
		Category:    "Sport",
		Link:        "www.onet.pl",
		Description: "test news onet",
	}

	fmt.Println(firstNews)
	fmt.Println(firstNews.Category)
}
