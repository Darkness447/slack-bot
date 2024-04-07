package main

import (
	"fmt"
	"log"

	"github.com/Darkness447/slack-bot/chat"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("BOT IS SETUP")

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	chat.Run()

	// api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	// channelArr := []string{os.Getenv("CHANNEL_ID")}
	// fileArr := []string{"we.txt", "why.txt"}

	// for i := 0; i < len(fileArr); i++ {
	// 	params := slack.FileUploadParameters{
	// 		Channels: channelArr,
	// 		File:     fileArr[i],
	// 	}

	// 	file, err := api.UploadFile(params)

	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	fmt.Printf("Name %s", file.Name)
	// }
}
