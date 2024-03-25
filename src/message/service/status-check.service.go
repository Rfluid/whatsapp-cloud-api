package message_service

import (
	bootstrap_model "github.com/Rfluid/whatsapp/src/bootstrap/model"
	message_model "github.com/Rfluid/whatsapp/src/message/model"
)

// Sends hello_world template to receiverNumber.
func SentMessageStatusCheck(
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
