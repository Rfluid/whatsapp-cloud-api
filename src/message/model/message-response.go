package message_model

import common_model "github.com/Rfluid/whatsapp-cloud-api/src/common/model"

type MessageResponse struct {
	MessageStatus string `json:"message_status,omitempty"`
	common_model.Id
}
