package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"
)

const testData="2000-09-07"


func main(){
	location,_:=time.LoadLocation("Asia/Tokyo")
	var is_celebrating bool = false
	month:=time.Now().In(location).Month()
	day:=time.Now().In(location).Day()
	target,_:=time.ParseInLocation("2006-01-02",testData,location)
	if month==target.Month() && day==target.Day(){
		is_celebrating=true
	}
	if is_celebrating{
		// 一致する場合はline bot を起動。
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
}

