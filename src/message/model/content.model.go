package message_model

import message_type_model "github.com/Rfluid/whatsapp/src/message/model/type"

// The message real content seen by user.
type Content struct {
	Text        *message_type_model.Text     `json:"text,omitempty"`
	Reaction    *message_type_model.Reaction `json:"reaction,omitempty"`
	Image       *message_type_model.Media    `json:"image,omitempty"`
	Video       *message_type_model.Media    `json:"video,omitempty"`
	Document    *message_type_model.Media    `json:"document,omitempty"`
	Audio       *message_type_model.Media    `json:"audio,omitempty"`
	Sticker     *message_type_model.Media    `json:"sticker,omitempty"`
	Location    *message_type_model.Location `json:"location,omitempty"`
	Interactive interface{}                  `json:"interactive,omitempty"`
	Template    interface{}                  `json:"template,omitempty"`
}

// TODO: Add validate content method.
