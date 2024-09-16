package message_type_interactive_model

import (
	media_model "github.com/Rfluid/whatsapp-cloud-api/src/media/model"
	message_type_common_model "github.com/Rfluid/whatsapp-cloud-api/src/message/model/content-type/common"
)

// Described at
//
// https://developers.facebook.com/docs/whatsapp/on-premises/reference/messages#header-object
type Header struct {
	Type message_type_common_model.Type `json:"type"` // Type of message.

	Text     *message_type_common_model.TextData `json:"text,omitempty"`
	Image    *media_model.UseMedia               `json:"image,omitempty"`
	Video    *media_model.UseMedia               `json:"video,omitempty"`
	Document *media_model.UseMedia               `json:"document,omitempty"`
	Audio    *media_model.UseMedia               `json:"audio,omitempty"`
}
