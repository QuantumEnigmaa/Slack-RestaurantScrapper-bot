package scrapper

import (
	"strings"

	"github.com/gocolly/colly"
)

type Dish struct {
	DishName    string
	Description string
}

func Scrapper(domain string, address string) []Dish {
	var list_dish []Dish
	collector := colly.NewCollector(colly.AllowedDomains(domain))

	collector.OnHTML("div[id=menu-content]", func(e *colly.HTMLElement) {
		var list_name []string
		var list_description []string

		dishName := strings.Split(e.ChildText("div.row > div.col-md-6 > h4.item-head"), "\n")
		for index, name := range dishName {
			if index > 0 {
				if strings.Contains(name, "GOURMANDISE") {
					break
				}
				n := strings.TrimSpace(name)

				list_name = append(list_name, n)
			}
		}

		description := strings.Split(e.ChildText("div.row > div.col-md-6 > div.item-description"), "\n")
		for _, desc := range description {
			d := strings.TrimSpace(desc)

			if d != "" && !strings.Contains(d, "Attention") {
				list_description = append(list_description, d)
			}
		}

		for index, current_dish := range list_name {
			new_dish := Dish{
				DishName:    current_dish,
				Description: list_description[index],
			}

			list_dish = append(list_dish, new_dish)
		}
	})

	collector.Visit(address)
	return list_dish
}
