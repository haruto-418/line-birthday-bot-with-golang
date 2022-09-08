package utils

import (
	"log"
	"time"
)

func IsCelebrating(target string)bool{
	location,_:=time.LoadLocation("Asia/Tokyo")
	month:=time.Now().In(location).Month()
	day:=time.Now().In(location).Day()
	target_time,err:=time.ParseInLocation("2006-01-02",target,location)
	if err!=nil{
		log.Fatal(err)
	}
	var is_celebrating bool = false
	if month==target_time.Month() && day==target_time.Day(){
		is_celebrating=true
	}
	return is_celebrating
}