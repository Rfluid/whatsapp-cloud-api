package message

import "github.com/Rfluid/whatsapp-cloud-api/src/common"

type MarkAsReadPayload struct {
	MessageID string        `json:"message_id"`
	Status    SendingStatus `json:"status" validate:"required,sending_status"`

	common.MessagingProduct
}
