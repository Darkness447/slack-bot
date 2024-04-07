package chat

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/slack-io/slacker"
)

func Run() {

	e := godotenv.Load(".env")

	if e != nil {
		log.Fatalf("Error loading .env file")
	}

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	bot.AddCommand(&slacker.CommandDefinition{
		Command:     "My Yob is <year>",
		Description: "Yob calculator",
		Examples:    []string{"My Yob is 2024"},
		Handler: func(botCtx *slacker.CommandContext) {
			fmt.Println("verifing")
			year := botCtx.Request().Param("year")
			yob, err := strconv.Atoi(year)

			if err != nil {
				fmt.Println(err)
			}

			age := 2024 - yob
			r := fmt.Sprintf("age is %d", age)
			botCtx.Response().Reply(r)
			// botCtx.Response().Reply("hi!")
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)

	if err != nil {
		log.Fatal(err)
	}

}
