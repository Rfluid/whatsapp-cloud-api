package template_model

type TemplateSummaryResponse struct {
	TotalCount              *uint32 `json:"total_count,omitempty"`
	MessageTemplateCount    *uint32 `json:"message_template_count,omitempty"`
	MessageTemplateLimit    *uint32 `json:"message_template_limit,omitempty"`
	AreTranslationsComplete *bool   `json:"are_translations_complete,omitempty"`
}

type TemplateSummary string

var (
	TotalCount              TemplateSummary = "total_count"
	MessageTemplateCount    TemplateSummary = "message_template_count"
	MessageTemplateLimit    TemplateSummary = "message_template_limit"
	AreTranslationsComplete TemplateSummary = "are_translations_complete"
)
