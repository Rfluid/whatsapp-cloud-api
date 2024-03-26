package media_model

import common_model "github.com/Rfluid/whatsapp-cloud-api/src/common/model"

type Upload struct {
	File                          string `json:"file"` // Path to file.
	Type                          Type   `json:"type"`
	common_model.MessagingProduct        // Use SetDefault() method to set this property.
}
