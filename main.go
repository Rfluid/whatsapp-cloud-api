package main

import (
	"encoding/json"
	"fmt"
	"log"

	bootstrap_model "github.com/Rfluid/whatsapp/src/bootstrap/model"
	bootstrap_service "github.com/Rfluid/whatsapp/src/bootstrap/service"
	message_model "github.com/Rfluid/whatsapp/src/message/model"
	message_type_model "github.com/Rfluid/whatsapp/src/message/model/type"
	message_service "github.com/Rfluid/whatsapp/src/message/service"
	"github.com/joho/godotenv"
)

var Api bootstrap_model.CloudApi

func main() {
	// Provisory so we can test
	Api = bootstrap_service.GenerateCloudAPI()
	loadEnv()

	msg := message_model.Message{
		Way: message_model.Way{
			To:   "",
			Type: message_model.Text,
		},
		Content: message_model.Content{
			Text: &message_type_model.Text{
				PreviewURL: false,
				Body:       "Usando a api oficial do zap.",
			},
		},
	}
	msg.SetDefault()

	var arr = make([]message_model.Message, 100)

	for i := range arr {
		arr[i] = msg
	}

	resp, err := message_service.SafeSpamMany(Api, arr, 3)

	jsonData, _ := json.Marshal(resp)
	fmt.Println(string(jsonData))

	fmt.Println(err)
}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	bootstrap_service.BootstrapApiValues(&Api)
}
