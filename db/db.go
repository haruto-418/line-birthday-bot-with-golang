package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func ConnectDb()*sql.DB{
	if err:=godotenv.Load(".env"); err!=nil{
		log.Fatal(err)
	}
	path:=fmt.Sprintf( "%s:%s@/%s?multiStatements=true",os.Getenv("MYSQL_USER_NAME"),os.Getenv("MYSQL_PASSWORD"),os.Getenv("MYSQL_DB_NAME"))
	db,err:=sql.Open("mysql",path)
	if err!=nil{
		log.Fatal(err)
	}
	return db
}