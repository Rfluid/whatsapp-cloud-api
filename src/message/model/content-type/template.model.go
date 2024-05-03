// Handles content validation.
package message_content_type_model

import (
	common_model "github.com/Rfluid/whatsapp-cloud-api/src/common/model"
	message_type_template_model "github.com/Rfluid/whatsapp-cloud-api/src/message/model/content-type/template"
)

type Template struct {
	Name       string                                  `json:"name"`
	Language   common_model.Language                   `json:"language"`
	Components []message_type_template_model.Component `json:"components"`
}
