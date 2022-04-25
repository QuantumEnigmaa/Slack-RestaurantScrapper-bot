package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

type dish struct {
	dishName    string
	description string
	price       int
}

func main() {
	scrapper("boutique.wysifood.fr", "https://boutique.wysifood.fr/5bde4fa022f9e550534f93537e604529/take-away/menu/15809")
}

func scrapper(domain string, address string) {
	collector := colly.NewCollector(colly.AllowedDomains(domain))

	collector.OnHTML("div[id=menu-content]", func(e *colly.HTMLElement) {
		/* e.ForEach(".row", func(_ int, el *colly.HTMLElement) {
			fmt.Printf("Dish name : %v\n", e.ChildText("div.col-md-6 > h4.item-head"))
		}) */

		dishName := strings.Split(e.ChildText("div.row > div.col-md-6 > h4.item-head"), "\n")
		for _, name := range dishName {
			strings.Fields(name)
			fmt.Printf("Dish Name : %v\n", strings.Fields(name))
		}

		fmt.Println("#####################")
		fmt.Println("DISH DESCRIPTION")

		description := strings.Split(e.ChildText("div.row > div.col-md-6 > div.item-description"), "\n")
		for _, desc := range description {
			strings.Fields(desc)
			fmt.Printf("Dish description : %v\n", strings.Fields(desc))
		}

		fmt.Println("#####################")
		fmt.Println("DISH PRICE")

		prices := strings.Split(e.ChildText("div.row > div.item-cart > span.item-cart-price"), "\n")
		for _, price := range prices {
			strings.Fields(price)
			fmt.Printf("Price : %v\n", strings.Fields(price))
		}
	})

	collector.Visit(address)
}
