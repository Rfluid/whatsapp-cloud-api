package validators

import (
	"strings"

	"github.com/Rfluid/whatsapp-cloud-api/src/template"
	"github.com/go-playground/validator/v10"
)

// templateCategoryValidation validates TemplateCategory values case-insensitively.
func templateCategoryValidation(fl validator.FieldLevel) bool {
	input := fl.Field().String()

	tc := template.TemplateCategory(strings.ToUpper(input))
	return tc.IsValid()
}

// RegisterTemplateCategoryValidator registers the "template_category" validator.
func RegisterTemplateCategoryValidator(v *validator.Validate) error {
	return v.RegisterValidation("template_category", templateCategoryValidation)
}
