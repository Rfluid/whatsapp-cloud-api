package message_model

import common_model "github.com/Rfluid/whatsapp-cloud-api/src/common/model"

type MessageResponse struct {
	MessageStatus SendingStatus `json:"message_status" validate:"required,sending_status"`
	common_model.Id
}
