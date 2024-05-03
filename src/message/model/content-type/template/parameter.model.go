package message_type_template_model

import (
	media_model "github.com/Rfluid/whatsapp-cloud-api/src/media/model"
	message_type_common_model "github.com/Rfluid/whatsapp-cloud-api/src/message/model/content-type/common"
)

type Parameter struct {
	Type     ParameterType                       `json:"type"`
	Image    *media_model.UseMedia               `json:"image,omitempty"`
	Video    *media_model.UseMedia               `json:"video,omitempty"`
	Document *media_model.UseMedia               `json:"document,omitempty"`
	Audio    *media_model.UseMedia               `json:"audio,omitempty"`
	Sticker  *media_model.UseMedia               `json:"sticker,omitempty"`
	Text     *message_type_common_model.Text     `json:"text,omitempty"`
	DateTime *message_type_common_model.DateTime `json:"date_time,omitempty"`
}
