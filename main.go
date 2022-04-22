package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	collector := colly.NewCollector(colly.AllowedDomains("boutique.wysifood.fr"))

	collector.OnHTML(".item-head", func(e *colly.HTMLElement) {
		fmt.Printf("Printing whole element : %v\n", e.Text)

	})

	collector.Visit("https://boutique.wysifood.fr/5bde4fa022f9e550534f93537e604529/take-away/menu")
}
