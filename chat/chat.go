package chat

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
	}
}

func Run() {
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_SOCKET_TOKEN"))
	go printCommandEvents(bot.CommandEvents())

	bot.Command("my birth  is <year>",&slacker.CommandDefinition{
		Description:"Yob calculator",
		Example:"My Yob is 2024",
		Handler: func(botCtx slacker.BotContext,request slacker.Request,response slacker.ResponseWriter){
			year := request.Param("year")
			yob ,err := strconv.Atoi(year)
			
			if err!=nil {
				fmt.Println(err)
			}

			age := 2024 - yob;
			
		}
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)

	if err != nil {
		log.Fatal(err)
	}


}
