// Provides models to handle messages.
package message_model

import (
	media_model "github.com/Rfluid/whatsapp-cloud-api/src/media/model"
	message_type_common_model "github.com/Rfluid/whatsapp-cloud-api/src/message/model/content-type/common"
)

// The message real content seen by user.
type Content struct {
	Text        *message_type_common_model.Text     `json:"text,omitempty"`
	Reaction    *message_type_common_model.Reaction `json:"reaction,omitempty"`
	Image       *media_model.UseMedia               `json:"image,omitempty"`
	Video       *media_model.UseMedia               `json:"video,omitempty"`
	Document    *media_model.UseMedia               `json:"document,omitempty"`
	Audio       *media_model.UseMedia               `json:"audio,omitempty"`
	Sticker     *media_model.UseMedia               `json:"sticker,omitempty"`
	Location    *message_type_common_model.Location `json:"location,omitempty"`
	Interactive interface{}                         `json:"interactive,omitempty"`
	Template    interface{}                         `json:"template,omitempty"`
}

// TODO: Add validate content method.
