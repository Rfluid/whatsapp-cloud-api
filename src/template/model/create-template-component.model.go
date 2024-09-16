package template_model

import (
	media_model "github.com/Rfluid/whatsapp-cloud-api/src/media/model"
	message_type_common_model "github.com/Rfluid/whatsapp-cloud-api/src/message/model/content-type/common"
)

type CreateTemplateComponent struct {
	Type     ComponentType                           `json:"type"`
	Text     string                                  `json:"text,omitempty"`
	Image    *media_model.UseMedia                   `json:"image,omitempty"`
	Video    *media_model.UseMedia                   `json:"video,omitempty"`
	Document *media_model.UseMedia                   `json:"document,omitempty"`
	Audio    *media_model.UseMedia                   `json:"audio,omitempty"`
	Sticker  *media_model.UseMedia                   `json:"sticker,omitempty"`
	DateTime *message_type_common_model.DateTime     `json:"date_time,omitempty"`
	Currency *message_type_common_model.Currency     `json:"currency,omitempty"`
	Buttons  *[]message_type_common_model.ButtonData `json:"buttons,omitempty"`

	Example *ComponentExample `json:"example"`
}
