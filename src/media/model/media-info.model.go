// Provides models to handle media.
package media_model

import common_model "github.com/Rfluid/whatsapp-cloud-api/src/common/model"

type MediaInfo struct {
	MessagingProduct string `json:"messaging_product,omitempty"`
	URL              string `json:"url,omitempty"`
	MimeType         string `json:"mime_type,omitempty"`
	Sha256           string `json:"sha256,omitempty"`
	FileSize         string `json:"file_size,omitempty"`
	common_model.Id
}
