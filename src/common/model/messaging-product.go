package common_model

type MessagingProduct struct {
	MessagingProduct string `json:"messaging_product"` // Default is "whatsapp".
}

func (mP *MessagingProduct) SetDefault() *MessagingProduct {
	mP.MessagingProduct = "whatsapp"
	return mP
}
