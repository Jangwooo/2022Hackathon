package main

import (
	"log"

	"github.com/Jangwooo/2022Hackathon/interner/controller"
	"github.com/Jangwooo/2022Hackathon/interner/mysql"
	"github.com/joho/godotenv"
)

var envRoot = "/Users/jwmbp/Documents/2022Hackathon/config/.env"

func main() {
	if err := godotenv.Load(envRoot); err != nil {
		log.Fatalf("critical error: %s", err.Error())
	}

	if err := mysql.Migration(); err != nil {
		log.Fatalf("critical error: %v", err)
	}

	r := controller.SetUp()

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("critical error: %s", err.Error())
	}
}
