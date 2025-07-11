package validators

import (
	"strings"

	template_model "github.com/Rfluid/whatsapp-cloud-api/src/template/model"
	"github.com/go-playground/validator/v10"
)

// templateComponentTypeValidation validates ComponentType values case-insensitively.
func templateComponentTypeValidation(fl validator.FieldLevel) bool {
	input, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	ct := template_model.ComponentType(strings.ToUpper(input))
	return ct.IsValid()
}

// RegisterComponentTypeValidator registers the "template_component_type" validator.
func RegisterTemplateComponentTypeValidator(v *validator.Validate) error {
	return v.RegisterValidation("template_component_type", templateComponentTypeValidation)
}
