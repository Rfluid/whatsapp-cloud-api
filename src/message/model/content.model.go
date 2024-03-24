package message_model

import message_content_type_model "github.com/Rfluid/whatsapp/src/message/model/content-type"

// The message real content seen by user.
type Content struct {
	Text        *message_content_type_model.Text     `json:"text,omitempty"`
	Reaction    *message_content_type_model.Reaction `json:"reaction,omitempty"`
	Image       *message_content_type_model.Media    `json:"image,omitempty"`
	Video       *message_content_type_model.Media    `json:"video,omitempty"`
	Document    *message_content_type_model.Media    `json:"document,omitempty"`
	Audio       *message_content_type_model.Media    `json:"audio,omitempty"`
	Sticker     *message_content_type_model.Media    `json:"sticker,omitempty"`
	Location    *message_content_type_model.Location `json:"location,omitempty"`
	Interactive interface{}                          `json:"interactive,omitempty"`
	Template    interface{}                          `json:"template,omitempty"`
}

// TODO: Add validate content method.
