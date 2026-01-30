package message

import "github.com/Rfluid/whatsapp-cloud-api/src/common"

type MessageResponse struct {
	MessageStatus SendingStatus `json:"message_status" validate:"required,sending_status"`
	common.ID
}
