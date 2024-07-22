package config

import (
	"log"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnetDB()	{
	db, err := sql.Open("mysql", "root:r0ngs0X@@/gowebnative?parseTime=true")
	if err != nil {
		panic(err)
	}

	log.Println("Connected to database")
	DB = db
}