package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"
)

func main(){
	if err:=godotenv.Load(".env"); err!=nil{
		log.Fatal(err)
	}
	bot,err:=linebot.New(
		os.Getenv("LINE_BOT_CHANNEL_SECRET"),
		os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
	)
	if err!=nil{
		log.Fatal(err)
	}
	message:=linebot.NewTextMessage("hello")
	if _,err:=bot.BroadcastMessage(message).Do(); err!=nil{
		log.Fatal(err)
	}
}