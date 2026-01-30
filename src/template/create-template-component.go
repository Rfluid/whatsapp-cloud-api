package template

import (
	"github.com/Rfluid/whatsapp-cloud-api/src/media"
	"github.com/Rfluid/whatsapp-cloud-api/src/message/content"
)

type TemplateComponent = CreateTemplateComponent

type CreateTemplateComponent struct {
	Type     ComponentType                       `json:"type" validate:"required,template_component_type"`
	Format   ComponentFormat                     `json:"format,omitempty"`
	Text     string                              `json:"text,omitempty"`
	Image    *media.UseMedia               `json:"image,omitempty"`
	Video    *media.UseMedia               `json:"video,omitempty"`
	Document *media.UseMedia               `json:"document,omitempty"`
	Audio    *media.UseMedia               `json:"audio,omitempty"`
	Sticker  *media.UseMedia               `json:"sticker,omitempty"`
	DateTime *content.DateTime `json:"date_time,omitempty"`
	Currency *content.Currency `json:"currency,omitempty"`
	Buttons  *[]ButtonData                       `json:"buttons,omitempty"`

	Example *ComponentExample `json:"example,omitempty"`
}
