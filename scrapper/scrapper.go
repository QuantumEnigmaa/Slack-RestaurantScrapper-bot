package scrapper

import (
	"encoding/csv"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

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
		re := regexp.MustCompile("^([A-Z]{2})+")
		for index, name := range dishName {
			if index > 0 {
				n := strings.TrimSpace(name)

				if re.MatchString(n) {
					break
				} else {
					list_name = append(list_name, n)
				}
			}
		}

		description := strings.Split(e.ChildText("div.row > div.col-md-6 > div.item-description"), "\n")
		for index, desc := range description {
			d := strings.TrimSpace(desc)

			if d != "" && !strings.Contains(d, "Attention") {
				if index > 1 && index < (len(description)-1) && strings.TrimSpace(description[index-1]) != "" && strings.TrimSpace(description[index+1]) == "" {
					for i, element := range list_description {
						if element == strings.TrimSpace(description[index-1]) {
							list_description[i] = strings.Join([]string{element, d}, " ")
							continue
						}
					}
				}
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

	if len(list_dish) > 0 {
		var menu [][]string
		for _, d := range list_dish {
			var dish []string

			dish = append(dish, d.DishName, time.Now().UTC().String())
			menu = append(menu, dish)
		}

		writeMenuToCsv("menu.csv", menu)
	}

	return list_dish
}

func writeMenuToCsv(filename string, data [][]string) {
	file, err := os.OpenFile(filename, os.O_RDWR, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.WriteAll(data)
}
