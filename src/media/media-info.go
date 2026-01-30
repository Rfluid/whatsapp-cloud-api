// Provides models to handle media.
package media

import "github.com/Rfluid/whatsapp-cloud-api/src/common"

type MediaInfo struct {
	MessagingProduct string `json:"messaging_product,omitempty"`
	URL              string `json:"url,omitempty"`
	MimeType         string `json:"mime_type,omitempty"`
	Sha256           string `json:"sha256,omitempty"`
	FileSize         int64  `json:"file_size,omitempty"`
	common.ID
}
