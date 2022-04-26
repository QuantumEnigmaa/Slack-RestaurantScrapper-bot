package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"restaurant-scrapper/scrapper"

	"github.com/shomali11/slacker"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-2149077173828-3465010216224-AhZoXj0rcGqZfS5061BiTuHP")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A03CZ6E5LLT-3441211155506-a9c380773f22ab50913c4255f3d9b97829ce66222f1983a7cb72264c7cb80f18")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("menu", &slacker.CommandDefinition{
		Description: "Spok menu descriptor",
		Example:     "menu",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			dishes := scrapper.Scrapper("boutique.wysifood.fr", "https://boutique.wysifood.fr/5bde4fa022f9e550534f93537e604529/take-away/menu/15809")

			var r string
			for _, d := range dishes {
				r = r + fmt.Sprintf("%s\n%s\n\n", d.DishName, d.Description)
			}

			response.Reply(r)
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
