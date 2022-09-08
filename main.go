package main

import (
	"log"
	"os"
	"time"

	db "github.com/haruto-418/line-birthday-bot-with-golang/database"
	"github.com/haruto-418/line-birthday-bot-with-golang/database/models"
	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"
)

const testData="2000-09-07"


func main(){
	db:=db.ConnectDb()
	mysql_db,_:=db.DB()
	defer mysql_db.Close()
	users:= []models.User{}
	t:=time.Now()
	db.Where("birthday = ?",t.Format("2006-01-02")).Find(&users)
	if len(users)==0{
		// 誕生日の人はいないから何も実行しない。
	}else{
		// 誕生日の人がいるから実行。	
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

