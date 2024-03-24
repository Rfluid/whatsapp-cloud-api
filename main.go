package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

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

	var wg sync.WaitGroup

	msg := message_model.Message{
		Way: message_model.Way{
			To:   "5591984288778",
			Type: message_model.Text,
		},
		Content: message_model.Content{
			Text: &message_type_model.Text{
				PreviewURL: false,
				Body:       "TOP G",
			},
		},
	}
	msg.SetDefault()

	fmt.Println(msg)

	message_service.SendMany(
		Api,
		[]message_model.Message{msg, msg, msg},
	)

	jsonData, _ := json.Marshal(msg)
	fmt.Println(string(jsonData))
	// message_service.SendInitialTemplate(Api)

	wg.Wait()
}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	bootstrap_service.BootstrapApiValues(&Api)
}
