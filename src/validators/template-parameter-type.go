package validators

import (
	"strings"

	template_model "github.com/Rfluid/whatsapp-cloud-api/src/template/model"
	"github.com/go-playground/validator/v10"
)

// templateParameterTypeValidation validates if a string is a valid ParameterType (case-insensitive).
func templateParameterTypeValidation(fl validator.FieldLevel) bool {
	input, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	pt := template_model.ParameterType(strings.ToLower(input))
	return pt.IsValid()
}

// RegisterTemplateParameterTypeValidator registers the "parameter_type" validator.
func RegisterTemplateParameterTypeValidator(v *validator.Validate) error {
	return v.RegisterValidation("template_parameter_type", templateParameterTypeValidation)
}
