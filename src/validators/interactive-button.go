package validators

import (
	"strings"

	"github.com/Rfluid/whatsapp-cloud-api/src/message/interactive"
	"github.com/go-playground/validator/v10"
)

// interactiveButtonTypeValidation is the actual validation logic for ButtonType (case-insensitive).
func interactiveButtonTypeValidation(fl validator.FieldLevel) bool {
	input := fl.Field().String()

	bt := interactive.ButtonType(strings.ToUpper(input))
	return bt.IsValid()
}

// RegisterTemplateButtonTypeValidator registers the custom "interactive_button_type" validator.
func RegisterInteractiveButtonTypeValidator(v *validator.Validate) error {
	return v.RegisterValidation("interactive_button_type", interactiveButtonTypeValidation)
}
