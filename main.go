package main

import (
	"fmt"
	"log"
	"os"
	"time"

	db "github.com/haruto-418/line-birthday-bot-with-golang/database"
	"github.com/haruto-418/line-birthday-bot-with-golang/database/models"
	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"
)



func main(){
	db:=db.ConnectDb()
	mysql_db,_:=db.DB()
	defer mysql_db.Close()
	users:= []models.User{}
	t:=time.Now()
	db.Where("birthday = ?",t.Format("2006-01-02")).Find(&users)
	if len(users)==0{
		fmt.Println("no one is celebrating.")
		// 誕生日の人はいないから何も実行しない。
	}else{
		fmt.Println(users)
		// 誕生日の人がいるなら実行。	
		if err:=godotenv.Load(".env"); err!=nil{
			log.Fatal(err)
		}
		bot,bot_err:=linebot.New(
			os.Getenv("LINE_BOT_CHANNEL_SECRET"),
			os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
		)
		if bot_err!=nil{
			log.Fatal(bot_err)
		}
		message:=linebot.NewTextMessage("hello")
		if _,err:=bot.BroadcastMessage(message).Do(); err!=nil{
			log.Fatal(err)
		}
	}
	
}

