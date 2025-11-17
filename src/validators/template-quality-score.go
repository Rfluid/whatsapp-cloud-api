package validators

import (
	"strings"

	template_model "github.com/Rfluid/whatsapp-cloud-api/src/template/model"
	"github.com/go-playground/validator/v10"
)

// templateQualityScoreValidation checks if the input string is a valid TemplateQualityScore value (case-insensitive).
func templateQualityScoreValidation(fl validator.FieldLevel) bool {
	input := fl.Field().String()

	templateQualityScore := template_model.TemplateQualityScore(strings.ToUpper(input))
	return templateQualityScore.IsValid()
}

// RegisterTemplateQualityScoreValidator registers the "template_quality_score" validator.
func RegisterTemplateQualityScoreValidator(v *validator.Validate) error {
	return v.RegisterValidation("template_quality_score", templateQualityScoreValidation)
}
