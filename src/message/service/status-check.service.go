package message_service

import (
	bootstrap_model "github.com/Rfluid/whatsapp-cloud-api/src/bootstrap/model"
	common_model "github.com/Rfluid/whatsapp-cloud-api/src/common/model"
	message_model "github.com/Rfluid/whatsapp-cloud-api/src/message/model"
	message_type_common_model "github.com/Rfluid/whatsapp-cloud-api/src/message/model/content-type/common"
	template_model "github.com/Rfluid/whatsapp-cloud-api/src/template/model"
)

// Checks API health sending hello_world template to receiverNumber.
func SendMessageStatusCheck(
	api bootstrap_model.WhatsAppAPI,
	receiverNumber string,
) (message_model.Response, error) {
	msg := message_model.Message{
		Direction: message_model.Direction{
			To:   receiverNumber,
			Type: message_type_common_model.Template,
		},
		Content: message_model.Content{
			Template: &template_model.UseTemplate{
				Name:     "hello_world",
				Language: common_model.Language{Code: "en_US"},
			},
		},
	}
	msg.SetDefault()

	req, err := Send(api, msg)

	return req, err
}
