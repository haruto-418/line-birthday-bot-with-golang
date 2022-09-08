package controllers

import (
	"log"

	"github.com/haruto-418/line-birthday-bot-with-golang/db/models"
	"gorm.io/gorm"
)


func GetUsers(db *gorm.DB)*gorm.DB{
	var users []models.User
	result:=db.Find(&users)
	if result.Error!=nil{
		log.Fatal(result.Error)
	}
	return result
}