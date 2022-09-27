package utils

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

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

func GenerateMessage(users string)string{
	messages:=[]string{fmt.Sprintf("今日は%sの誕生日です！みなさん心してお祝いしましょう！",users),fmt.Sprintf("今日という良き日は%sの誕生日です。盛大な祝福を捧げましょう。",users),fmt.Sprintf("今日は%sの誕生日だドン！みんなでお祝いするドン！",users)}
	rand.Seed(time.Now().Unix())
	return messages[rand.Intn(len(messages))]
}