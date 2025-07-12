package validators

import (
	"strings"

	message_type_interactive_model "github.com/Rfluid/whatsapp-cloud-api/src/message/model/content-type/interactive"
	"github.com/go-playground/validator/v10"
)

// interactiveButtonTypeValidation is the actual validation logic for ButtonType (case-insensitive).
func interactiveButtonTypeValidation(fl validator.FieldLevel) bool {
	input := fl.Field().String()

	bt := message_type_interactive_model.ButtonType(strings.ToUpper(input))
	return bt.IsValid()
}

// RegisterTemplateButtonTypeValidator registers the custom "interactive_button_type" validator.
func RegisterInteractiveButtonTypeValidator(v *validator.Validate) error {
	return v.RegisterValidation("interactive_button_type", interactiveButtonTypeValidation)
}
