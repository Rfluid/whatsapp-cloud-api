package main

import (
	"fmt"
	"log"

	bootstrap_model "github.com/Rfluid/whatsapp/src/bootstrap/model"
	bootstrap_service "github.com/Rfluid/whatsapp/src/bootstrap/service"
	"github.com/joho/godotenv"
)

var Api bootstrap_model.ApiBootstrap = bootstrap_model.ApiBootstrap{
	Prefix:  "https://graph.facebook.com",
	Version: "v18.0",
}

func main() {
	loadEnv()
}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	bootstrap_service.BootstrapApiValues(&Api)

	fmt.Println(Api.MainURL)
}
