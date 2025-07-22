package main

import (
	"database/sql"
	"fmt"
	"log"
//	"github.com/gofiber/fiber/v2"
  "github.com/go-sql-driver/mysql"
  "github.com/joho/godotenv"
)


type Car struct {
	ID        int    `json:"id"`
	Brand     string `json:"brand"`
	Model    string `json:"model"`
	Year      int    `json:"year"`
}

var db *sql.DB

func main()  {
envFile, _ := godotenv.Read(".env")
fmt.Println(envFile)
    cfg := mysql.NewConfig()
		cfg.User = envFile["DB_USER"] 
		cfg.Passwd = envFile["DB_PASS"]
    cfg.Net = "tcp"
    cfg.Addr = "127.0.0.1:3306"
    cfg.DBName = "carsdb"
  var err error

fmt.Printf("DSN: %q\n", cfg)
  db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
        log.Fatal(err)
    }
	pingErr := db.Ping()
  if pingErr != nil {
		log.Fatal(pingErr)
    }
  fmt.Println("Connected!")
}

defer db.Close()
