package validators

import (
	"strings"

	message_type_common_model "github.com/Rfluid/whatsapp-cloud-api/src/message/model/content-type/common"
	"github.com/go-playground/validator/v10"
)

// templateButtonSubTypeValidation validates if a string is a valid ButtonSubType (case-insensitive).
func templateButtonSubTypeValidation(fl validator.FieldLevel) bool {
	input, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	bst := message_type_common_model.ButtonSubType(strings.ToLower(input))
	return bst.IsValid()
}

// RegisterTemplateButtonSubTypeValidator registers the "template_button_sub_type" validator.
func RegisterTemplateButtonSubTypeValidator(v *validator.Validate) error {
	return v.RegisterValidation("template_button_sub_type", templateButtonSubTypeValidation)
}
