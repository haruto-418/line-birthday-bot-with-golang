package main

import (
	"fmt"
	"log"
	"time"

	db "github.com/haruto-418/line-birthday-bot-with-golang/database"
	"github.com/haruto-418/line-birthday-bot-with-golang/database/models"
	"github.com/haruto-418/line-birthday-bot-with-golang/utils"
	"github.com/line/line-bot-sdk-go/linebot"
)



func main(){
	db:=db.ConnectDb()
	mysql_db,_:=db.DB()
	defer mysql_db.Close()
	users:= []models.User{}
	// t:=time.Now()
	t,_:=time.Parse("2006-01-02","2022-09-09")
	db.Where("birthday = ?",t.Format("2006-01-02")).Find(&users)
	if len(users)==0{
		fmt.Println("no one is celebrating.")
		// 誕生日の人はいないから何も実行しない。
	}else{
		fmt.Println(users)
		for _,user:=range users{
			fmt.Println(user.UserName)
		}
		// 誕生日の人がいるなら実行。	
		bot:=utils.InitBot()
		message:=linebot.NewTextMessage("hello")
		if _,err:=bot.BroadcastMessage(message).Do(); err!=nil{
			log.Fatal(err)
		}
	}
	
}

