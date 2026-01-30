package validators

import (
	"strings"

	"github.com/Rfluid/whatsapp-cloud-api/src/template"
	"github.com/go-playground/validator/v10"
)

// templateParameterTypeValidation validates if a string is a valid ParameterType (case-insensitive).
func templateParameterTypeValidation(fl validator.FieldLevel) bool {
	input := fl.Field().String()

	pt := template.ParameterType(strings.ToLower(input))
	return pt.IsValid()
}

// RegisterTemplateParameterTypeValidator registers the "parameter_type" validator.
func RegisterTemplateParameterTypeValidator(v *validator.Validate) error {
	return v.RegisterValidation("template_parameter_type", templateParameterTypeValidation)
}
