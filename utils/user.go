package utils

import (
	"crypto/rand"
	"log"

	"github.com/haruto-418/line-birthday-bot-with-golang/database/models"
	"gorm.io/gorm"
)

func CreateRandomUser(db *gorm.DB,count uint8){
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	mysql_db,_:=db.DB()
	defer mysql_db.Close()
	var i uint8=0
	for i<count{
		b:=make([]uint8,count)
		if _,err:=rand.Read(b); err!=nil{
			log.Fatal(err)
		}
		var user_name string
		for _,v:=range b{
			user_name+=string(letters[int(v)%len(letters)])
		}
		random_date:=GenerateRandomDate()
		user:=models.User{UserName: user_name,Birthday: random_date}
		db.Create(&user)
		i++
	}
}

// func DeleteAllUsers(db *gorm.DB){
// 	mysql_db,_:=db.DB()
// 	defer mysql_db.Close()
// 	user:=models.User{}
// 	db.Delete(&user)
// }