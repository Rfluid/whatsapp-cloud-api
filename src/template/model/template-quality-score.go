package template_model

type TemplateQualityScore string

const (
	Green   TemplateQualityScore = "GREEN"
	Yellow  TemplateQualityScore = "YELLOW"
	Red     TemplateQualityScore = "RED"
	Unknown TemplateQualityScore = "UNKNOWN"
)

// IsValid checks if the quality score is one of the predefined values.
func (tqs TemplateQualityScore) IsValid() bool {
	switch tqs {
	case Green, Yellow, Red, Unknown:
		return true
	default:
		return false
	}
}
