package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"
)

func InitBot()*linebot.Client{
	if err:=godotenv.Load(".env"); err!=nil{
		log.Fatal(err)
	}
	bot,bot_err:=linebot.New(	os.Getenv("LINE_BOT_CHANNEL_SECRET"),os.Getenv("LINE_BOT_CHANNEL_TOKEN"),)
	if bot_err!=nil{
		log.Fatal(bot_err)
	}
	return bot
}