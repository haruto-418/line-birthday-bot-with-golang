package main

import (
	"fmt"
	"log"
	"os"

	"github.com/haruto-418/line-birthday-bot-with-golang/db"
	"github.com/haruto-418/line-birthday-bot-with-golang/db/controllers"
	"github.com/haruto-418/line-birthday-bot-with-golang/utils"
	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"
)

const testData="2000-09-07"


func main(){
	is_celebrating:=utils.IsCelebrating(testData)
	db:=db.ConnectDb()
	users:=controllers.GetUsers(db)
	fmt.Println(users)
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

