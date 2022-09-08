package models

import "time"


type User struct{
	ID 				int			 `gorm:"primaryKey"`
	UserName 	string
	Birthday  time.Time
}