package message_model

import common_model "github.com/Rfluid/whatsapp/src/common/model"

// Default fields for messaging.
type Default struct {
	RecipientType string `json:"recipient_type"` // Default is "individual"
	common_model.MessagingProduct
}

// Sets the default fields so you don't have to worry about it.
func (d *Default) SetDefault() {
	d.MessagingProduct.SetDefault()
	d.RecipientType = "individual"
}
