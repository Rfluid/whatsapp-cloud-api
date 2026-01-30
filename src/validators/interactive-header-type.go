package validators

import (
	"strings"

	"github.com/Rfluid/whatsapp-cloud-api/src/message/interactive"
	"github.com/go-playground/validator/v10"
)

// interactiveHeaderTypeValidation checks if a string is a valid HeaderType (case-insensitive).
func interactiveHeaderTypeValidation(fl validator.FieldLevel) bool {
	input := fl.Field().String()

	ht := interactive.HeaderType(strings.ToLower(input))
	return ht.IsValid()
}

// RegisterInteractiveHeaderTypeValidator registers the "interactive_header_type" validator.
func RegisterInteractiveHeaderTypeValidator(v *validator.Validate) error {
	return v.RegisterValidation("interactive_header_type", interactiveHeaderTypeValidation)
}
