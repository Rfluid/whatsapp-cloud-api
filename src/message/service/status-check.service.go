package message_service

import (
	bootstrap_model "github.com/Rfluid/whatsapp-cloud-api/src/bootstrap/model"
	message_model "github.com/Rfluid/whatsapp-cloud-api/src/message/model"
)

// Checks API health sending hello_world template to receiverNumber.
func SendMessageStatusCheck(
	api bootstrap_model.CloudApi,
	receiverNumber string,
) (message_model.Response, error) {
	msg := message_model.Message{
		Direction: message_model.Direction{
			To:   receiverNumber,
			Type: message_model.Template,
		},
		Content: message_model.Content{
			Template: map[string]interface{}{
				"name": "hello_world",
				"language": map[string]string{
					"code": "en_US",
				},
			},
		},
	}
	msg.SetDefault()

	req, err := Send(api, msg)

	return req, err
}
