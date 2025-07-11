package message_type_interactive_model

import (
	message_type_common_model "github.com/Rfluid/whatsapp-cloud-api/src/message/model/content-type/common"
)

type HeaderType message_type_common_model.Type

const (
	Document HeaderType = HeaderType(message_type_common_model.Document)
	Image    HeaderType = HeaderType(message_type_common_model.Image)
	Video    HeaderType = HeaderType(message_type_common_model.Video)
	Text     HeaderType = HeaderType(message_type_common_model.Text)
)

// IsValid checks if the HeaderType is one of the allowed values.
func (ht HeaderType) IsValid() bool {
	switch ht {
	case Document, Image, Video, Text:
		return true
	default:
		return false
	}
}
