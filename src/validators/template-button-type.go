package validators

import (
	"strings"

	template_model "github.com/Rfluid/whatsapp-cloud-api/src/template/model"
	"github.com/go-playground/validator/v10"
)

// templateButtonTypeValidation is the actual validation logic for ButtonType (case-insensitive).
func templateButtonTypeValidation(fl validator.FieldLevel) bool {
	input := fl.Field().String()

	bt := template_model.ButtonType(strings.ToUpper(input))
	return bt.IsValid()
}

// RegisterTemplateButtonTypeValidator registers the custom "template_button_type" validator.
func RegisterTemplateButtonTypeValidator(v *validator.Validate) error {
	return v.RegisterValidation("template_button_type", templateButtonTypeValidation)
}
