package message

import (
	"github.com/Rfluid/whatsapp-cloud-api/src/bootstrap"
	"github.com/Rfluid/whatsapp-cloud-api/src/common"
	"github.com/Rfluid/whatsapp-cloud-api/src/message/content"
	"github.com/Rfluid/whatsapp-cloud-api/src/template"
)

// Checks API health sending hello_world template to receiverNumber.
func SendMessageStatusCheck(
	api bootstrap.WhatsAppAPI,
	receiverNumber string,
) (Response, error) {
	msg := Message{
		Direction: Direction{
			To:   receiverNumber,
			Type: content.Template,
		},
		Content: Content{
			Template: &template.UseTemplate{
				Name:     "hello_world",
				Language: common.Language{Code: "en_US"},
			},
		},
	}
	msg.SetDefault()

	req, err := Send(api, msg)

	return req, err
}
