package main

import (
	"fmt"
	"log"
	"strings"
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
	t:=time.Now()
	db.Where("birthday = ?",t.Format("2006-01-02")).Find(&users)
	if len(users)==0{
		fmt.Println("no one is celebrating.")
		// 誕生日の人はいないから何も実行しない。
	}else{
		var birthday_users []string
		for _,user:=range users{
			birthday_users=append(birthday_users, user.UserName)
		}
		var users_text string
		if len(birthday_users)>1{
			users_text=strings.Join(birthday_users,"氏と")
		}else{
			users_text+=users_text+"氏"
		}
		// 誕生日の人がいるなら実行。
		bot:=utils.InitBot()
		message_text:=utils.GenerateMessage(users_text)
		message:=linebot.NewTextMessage(message_text)
		if _,err:=bot.BroadcastMessage(message).Do(); err!=nil{
			log.Fatal(err)
		}
	}
}

