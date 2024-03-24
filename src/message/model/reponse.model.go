package message_model

// Response by api when message is sent.
type Response struct {
	MessagingProduct string            `json:"messaging_product"` // Default "whatsapp". Actually you will only receive this.
	Contacts         []ResponseContact `json:"contacts"`          // Contacts that received messages.
	Messages         []Id              `json:"messages"`          // Ids of sent messages.
}
