package message_model

import common_model "github.com/Rfluid/whatsapp/src/common/model"

// Body to update message status.
type MessageStatus struct {
	Status string `json:"status,omitempty"` // Set "read" to mark as read.
	Context
	common_model.MessagingProduct
}
