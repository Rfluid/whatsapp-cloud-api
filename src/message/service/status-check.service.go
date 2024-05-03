package message_service

import (
	bootstrap_model "github.com/Rfluid/whatsapp-cloud-api/src/bootstrap/model"
	common_model "github.com/Rfluid/whatsapp-cloud-api/src/common/model"
	message_model "github.com/Rfluid/whatsapp-cloud-api/src/message/model"
	message_content_type_model "github.com/Rfluid/whatsapp-cloud-api/src/message/model/content-type"
)

// Checks API health sending hello_world template to receiverNumber.
func SendMessageStatusCheck(
	api bootstrap_model.WhatsAppAPI,
	receiverNumber string,
) (message_model.Response, error) {
	msg := message_model.Message{
		Direction: message_model.Direction{
			To:   receiverNumber,
			Type: message_model.Template,
		},
		Content: message_model.Content{
			Template: &message_content_type_model.Template{
				Name:     "hello_world",
				Language: common_model.Language{Code: "en_US"},
			},
		},
	}
	msg.SetDefault()

	req, err := Send(api, msg)

	return req, err
}
