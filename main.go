package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"restaurant-scrapper/scrapper"

	"github.com/enescakir/emoji"
	"github.com/fatih/color"
	"github.com/shomali11/slacker"
	"github.com/slack-go/slack"
)

func main() {
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("menu-spok", &slacker.CommandDefinition{
		Description: "Spok menu descriptor",
		Example:     "menu-spok",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			dishes := scrapper.Scrapper("boutique.wysifood.fr", "https://boutique.wysifood.fr/5bde4fa022f9e550534f93537e604529/take-away/menu/15809")

			italic := color.New(color.Italic).SprintFunc()
			vege := regexp.MustCompile(`\(V\)|veggie|\(\ V\)|\(V\ \)|\(\ V\ \)`)

			r := "---------Plats principaux (généralement 11.90eu)---------\n"
			for _, d := range dishes {
				if vege.MatchString(d.DishName) {
					r = r + fmt.Sprintf("%v\n%s\n\n", d.DishName+" "+string(emoji.LeafyGreen), italic(d.Description))
				} else {
					r = r + fmt.Sprintf("%s\n%s\n\n", d.DishName, italic(d.Description))
				}
			}

			response.Reply(r)
		},
	})

	bot.Command("get-menu-file", &slacker.CommandDefinition{
		Description: "Upload on chat the csv file containing all past menus",
		Example:     "get-menu-file",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			ev := botCtx.Event()
			client := botCtx.Client()

			if ev.Channel != "" {
				_, err := client.UploadFile(slack.FileUploadParameters{File: "menu.csv", Channels: []string{ev.Channel}})
				if err != nil {
					fmt.Printf("Error encountered when uploading file: %+v\n", err)
				}
			}
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func printCommandEvents(analyticChannel <-chan *slacker.CommandEvent) {
	for event := range analyticChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}
