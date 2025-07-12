package validators

import (
	"strings"

	message_type_interactive_model "github.com/Rfluid/whatsapp-cloud-api/src/message/model/content-type/interactive"
	"github.com/go-playground/validator/v10"
)

// interactiveHeaderTypeValidation checks if a string is a valid HeaderType (case-insensitive).
func interactiveHeaderTypeValidation(fl validator.FieldLevel) bool {
	input := fl.Field().String()

	ht := message_type_interactive_model.HeaderType(strings.ToLower(input))
	return ht.IsValid()
}

// RegisterInteractiveHeaderTypeValidator registers the "interactive_header_type" validator.
func RegisterInteractiveHeaderTypeValidator(v *validator.Validate) error {
	return v.RegisterValidation("interactive_header_type", interactiveHeaderTypeValidation)
}
