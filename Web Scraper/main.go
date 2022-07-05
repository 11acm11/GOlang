package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"github.com/gocolly/colly"
)

var url string = "https://www.imdb.com/list/ls058654847/"

type Standings struct {
	pos    string
	name   string
	year   string
	genre  string
	rating string
}

func main() {
	Stand := scrape(url)
	file, err := os.Create("data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, v := range Stand {
		writer.Write([]string{
			v.pos,
			v.name,
			v.year,
			v.rating,
			v.genre,
		})
	}

}
func scrape(url string) []Standings {
	standings := Standings{}
	anime := make([]Standings, 0)
	c := colly.NewCollector()
	c.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL)
	})
	c.OnHTML(".lister-item-content", func(e *colly.HTMLElement) {
		standings.pos = e.ChildText("span.lister-item-index.unbold.text-primary")
		standings.year = e.ChildText("span.lister-item-year.text-muted.unbold")
		standings.genre = e.ChildText("span.genre")
		anime = append(anime, standings)
	})
	j := 0
	c.OnHTML(".ipl-rating-star.small", func(e *colly.HTMLElement) {
		anime[j].rating = e.ChildText("span")
		j = j + 1
	})
	i := 0
	c.OnHTML(".lister-item-header", func(e *colly.HTMLElement) {
		anime[i].name = e.ChildText("a")
		i = i + 1
	})
	c.Visit(url)
	return anime
}
