package main

import (
	"log"

	"github.com/Jangwooo/2022Hackathon/docs"
	"github.com/Jangwooo/2022Hackathon/interner/controller"
	"github.com/Jangwooo/2022Hackathon/interner/mysql"
	"github.com/joho/godotenv"
)

var envRoot = "./config/.env"

func main() {
	if err := godotenv.Load(envRoot); err != nil {
		log.Fatalf("critical error: %s", err.Error())
	}

	if err := mysql.Migration(); err != nil {
		log.Fatalf("critical error: %v", err)
	}

	docs.SwaggerInfo.Title = "손길 API Docs"
	docs.SwaggerInfo.Host = "15.165.88.215:8000"
	docs.SwaggerInfo.BasePath = "/"

	r := controller.SetUp()

	if err := r.Run(":8000"); err != nil {
		log.Fatalf("critical error: %s", err.Error())
	}
}
