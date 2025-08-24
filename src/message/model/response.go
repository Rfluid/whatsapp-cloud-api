package message_model

import common_model "github.com/Rfluid/whatsapp-cloud-api/src/common/model"

// Response by api when message is sent.
type Response struct {
	Contacts []ResponseContact `json:"contacts,omitempty"` // Contacts that received messages.
	Messages []MessageResponse `json:"messages,omitempty"` // IDs of sent messages.
	common_model.MessagingProduct
}
