package message

import "github.com/Rfluid/whatsapp-cloud-api/src/common"

// Response by api when message is sent.
type Response struct {
	Contacts []ResponseContact `json:"contacts,omitempty"` // Contacts that received messages.
	Messages []MessageResponse `json:"messages,omitempty"` // IDs of sent messages.
	common.MessagingProduct
}
