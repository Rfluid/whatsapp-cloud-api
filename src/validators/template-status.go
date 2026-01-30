package validators

import (
	"strings"

	"github.com/Rfluid/whatsapp-cloud-api/src/template"
	"github.com/go-playground/validator/v10"
)

// templateStatusValidation checks if the input string is a valid TemplateStatus value (case-insensitive).
func templateStatusValidation(fl validator.FieldLevel) bool {
	input := fl.Field().String()

	templateStatus := template.Status(strings.ToUpper(input))
	return templateStatus.IsValid()
}

// RegisterTemplateStatusValidator registers the "template_status" validator.
func RegisterTemplateStatusValidator(v *validator.Validate) error {
	return v.RegisterValidation("template_status", templateStatusValidation)
}
