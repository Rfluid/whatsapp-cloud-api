package message_model

// Response by api when message is sent.
type Response struct {
	MessagingProduct string            `json:"messaging_product,omitempty"` // Default "whatsapp". Actually you will only receive this.
	Contacts         []ResponseContact `json:"contacts,omitempty"`          // Contacts that received messages.
	Messages         []Id              `json:"messages,omitempty"`          // Ids of sent messages.
}
