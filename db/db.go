package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDb()*gorm.DB{
	if err:=godotenv.Load(".env"); err!=nil{
		log.Fatal(err)
	}
	dsn:=fmt.Sprintf("%s:%s@/%s?multiStatements=true",os.Getenv("MYSQL_USER_NAME"),os.Getenv("MYSQL_PASSWORD"),os.Getenv("MYSQL_DB_NAME"))
	db,err:=gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err!=nil{
		log.Fatal(err)
	}
	return db
}