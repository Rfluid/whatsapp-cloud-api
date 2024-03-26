package media_model

import (
	"net/url"

	common_model "github.com/Rfluid/whatsapp-cloud-api/src/common/model"
)

type Upload struct {
	File                          string                          `json:"file"` // Path to file.
	Type                          common_model.SupportedMimeTypes `json:"type"`
	common_model.MessagingProduct                                 // Use SetDefault() method to set this property.
}

func (u *Upload) ToURLValues() url.Values {
	formData := url.Values{}

	formData.Set("file", string(u.File))
	formData.Set("type", string(u.Type))
	formData.Set("messaging_product", u.MessagingProduct.MessagingProduct)

	return formData
}
