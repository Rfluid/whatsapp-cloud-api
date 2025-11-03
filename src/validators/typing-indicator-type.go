package validators

import (
	"strings"

	typing_model "github.com/Rfluid/whatsapp-cloud-api/src/typing/model"
	"github.com/go-playground/validator/v10"
)

// typingIndicatorTypeValidation checks if a string is a valid SendingStatus (case-insensitive).
func typingIndicatorTypeValidation(fl validator.FieldLevel) bool {
	input := fl.Field().String()

	status := typing_model.TypingIndicatorType(strings.ToLower(input))
	return status.IsValid()
}

// RegisterTypingIndicatorTypeValidator registers the "typing_indicator_type" validator.
func RegisterTypingIndicatorTypeValidator(v *validator.Validate) error {
	return v.RegisterValidation("typing_indicator_type", typingIndicatorTypeValidation)
}
