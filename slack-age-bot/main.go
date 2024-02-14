package main

import (
	"context"
	"fmt"
	"github.com/shomali11/slacker"
	"log"
	"os"
	"strconv"
	"time"
)

func printComandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()

	}
}
func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-6600582270721-6581383322390-sg3R7qDJj6UtwLKXvHmxTEWY")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A06H9QX5WD9-6581326234214-9d71d7cc9011e848a6a8f05058a2c6fd33e6b989e9ed06b713ca01c0b7726610")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printComandEvents(bot.CommandEvents())

	bot.Command("my job is <year>", &slacker.CommandDefinition{
		Description: "Yob calculator",
		BlockID:     "My yob is 2020",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)

			if err != nil {
				println("error")
			}

			age := time.Now().Year() - yob

			r := fmt.Sprintf("Age is %d", age)
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
