package template_model

import message_type_common_model "github.com/Rfluid/whatsapp-cloud-api/src/message/model/content-type/common"

type Component struct {
	Type       ComponentType `json:"type" validate:"required,template_component_type"`
	Parameters []Parameter   `json:"parameters"`

	SubType message_type_common_model.ButtonSubType `json:"sub_type,omitempty" validate:"omitempty,template_button_sub_type"` // Only for button type.
	Index   string                                  `json:"index,omitempty"`                                                  // Only for button type. This is actually an integer string.
}
