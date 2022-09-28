package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	db "github.com/haruto-418/line-birthday-bot-with-golang/database"
	"github.com/haruto-418/line-birthday-bot-with-golang/database/models"
	"github.com/haruto-418/line-birthday-bot-with-golang/utils"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)



func main(){
	db:=db.ConnectDb()
	mysql_db,_:=db.DB()
	defer mysql_db.Close()
	users:= []models.User{}
	t:=time.Now()
	pattern:="%"+t.Format("01-02")
	db.Where("birthday LIKE ?",pattern).Find(&users)
	if len(users)==0{
		fmt.Println("no one is celebrating.")
		// 誕生日の人はいないから何も実行しない。
	}else{
		// 誕生日の人がいるなら実行。
		var birthday_users []string
		for _,user:=range users{
			birthday_users=append(birthday_users, user.UserName)
		}
		var users_text string
		if len(birthday_users)>1{
			users_text=strings.Join(birthday_users,"氏と")
		}else{
			users_text=birthday_users[0]+"氏"
		}
		bot:=utils.InitBot()
		message_text:=utils.GenerateMessage(users_text)
		message:=linebot.NewTextMessage(message_text)
		if _,err:=bot.BroadcastMessage(message).Do(); err!=nil{
			log.Fatal(err)
		}
	}
}

