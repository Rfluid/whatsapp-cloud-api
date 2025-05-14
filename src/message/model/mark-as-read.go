package message_model

import common_model "github.com/Rfluid/whatsapp-cloud-api/src/common/model"

type MarkAsRead struct {
	MessageId string        `json:"message_id"`
	Status    SendingStatus `json:"status"`

	common_model.MessagingProduct
}
