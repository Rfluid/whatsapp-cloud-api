package message_model

import common_model "github.com/Rfluid/whatsapp-cloud-api/src/common/model"

// Body to update message status.
type MessageStatus struct {
	Status SendingStatus `json:"status" validate:"required,sending_status"` // Set "read" to mark as read.
	Context
	common_model.MessagingProduct
}
