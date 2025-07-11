package template_model

type TemplateCategory string

const (
	Authentication TemplateCategory = "AUTHENTICATION"
	Marketing      TemplateCategory = "MARKETING"
	Utility        TemplateCategory = "UTILITY"
)

// IsValid checks if the category is one of the predefined ones.
func (tc TemplateCategory) IsValid() bool {
	switch tc {
	case Authentication, Marketing, Utility:
		return true
	default:
		return false
	}
}
