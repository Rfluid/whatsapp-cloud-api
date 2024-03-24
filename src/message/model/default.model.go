package message_model

// Default fields for messaging.
type Default struct {
	MessagingProduct string `json:"messaging_product"` // Default is "whatsapp"
	RecipientType    string `json:"recipient_type"`    // Default is "individual"
}

// Sets the default fields so you don't have to worry about it.
func (d *Default) SetDefault() {
	d.MessagingProduct = "whatsapp"
	d.RecipientType = "individual"
}
