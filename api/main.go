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
	articles := []Article{
		{
			Title:       "Test Onet",
			Source:      "Onet.pl",
			Category:    "Sport",
			Link:        "www.onet.pl",
			Description: "test news onet",
		},
		{
			Title:       "Test Wp",
			Source:      "Wp.pl",
			Category:    "Info",
			Link:        "www.wp.pl",
			Description: "test news wp",
		},
		{
			Title:       "Test o2",
			Source:      "o2.pl",
			Category:    "Gossip",
			Link:        "www.02.pl",
			Description: "test news o2",
		},
		{
			Title:       "Test gazeta",
			Source:      "gazeta.pl",
			Category:    "Moto",
			Link:        "www.gazeta.pl",
			Description: "test news gazeta",
		},
	}

	for _, s := range articles {
		fmt.Println(s)
	}

}
