package main

import (
	"encoding/json"
	"fmt"
	"log"

	bootstrap_model "github.com/Rfluid/whatsapp/src/bootstrap/model"
	bootstrap_service "github.com/Rfluid/whatsapp/src/bootstrap/service"
	message_model "github.com/Rfluid/whatsapp/src/message/model"
	message_content_type_model "github.com/Rfluid/whatsapp/src/message/model/content-type"
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
			To:   "5591984288778",
			Type: message_model.Text,
		},
		Content: message_model.Content{
			Text: &message_content_type_model.Text{
				PreviewURL: false,
				Body:       "because ** was in paris",
			},
		},
	}
	msg.SetDefault()

	resp, err := message_service.Send(Api, msg)

	jsonData, _ := json.Marshal(resp)
	fmt.Println(string(jsonData))

	jsonDataMsg, _ := json.Marshal(msg)
	fmt.Println(string(jsonDataMsg))

	fmt.Println(err)
}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	bootstrap_service.BootstrapApiValues(&Api)
}
