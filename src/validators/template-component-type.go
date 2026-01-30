package validators

import (
	"strings"

	"github.com/Rfluid/whatsapp-cloud-api/src/template"
	"github.com/go-playground/validator/v10"
)

// templateComponentTypeValidation validates ComponentType values case-insensitively.
func templateComponentTypeValidation(fl validator.FieldLevel) bool {
	input := fl.Field().String()

	ct := template.ComponentType(strings.ToUpper(input))
	return ct.IsValid()
}

// RegisterComponentTypeValidator registers the "template_component_type" validator.
func RegisterTemplateComponentTypeValidator(v *validator.Validate) error {
	return v.RegisterValidation("template_component_type", templateComponentTypeValidation)
}
