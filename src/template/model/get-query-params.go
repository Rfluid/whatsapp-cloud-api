package template_model

import (
	"net/url"
	"strconv"
	"strings"

	common_model "github.com/Rfluid/whatsapp-cloud-api/src/common/model"
)

type TemplateQueryParams struct {
	Name          string `json:"name,omitempty"`
	Content       string `json:"content,omitempty"`
	Language      string `json:"language,omitempty"`
	Status        Status `json:"status,omitempty" validate:"omitempty,template_status"`
	Category      string `json:"category,omitempty"`
	NameOrContent string `json:"name_or_content,omitempty" query:"name_or_content"`

	Limit *uint64 `json:"limit,omitempty" query:"limit"`

	Fields  *[]TemplateFields  `json:"fields" query:"fields"`   // Fields to be returned.
	Summary *[]TemplateSummary `json:"summary" query:"summary"` // Summary to be returned.

	common_model.GraphCursors
}

// BuildQuery constructs the query string from TemplateQueryParams.
// It returns an encoded query string that can be appended to a URL.
func (qp *TemplateQueryParams) BuildQuery() string {
	// Initialize a URL values object
	v := url.Values{}

	// Conditionally set each query parameter if it's provided

	if qp.Name != "" {
		v.Set("name", qp.Name)
	}

	if qp.Limit != nil {
		v.Set("limit", strconv.FormatUint(*qp.Limit, 10))
	}

	if qp.NameOrContent != "" {
		v.Set("name_or_content", qp.Name)
	}

	if qp.Content != "" {
		v.Set("components", qp.Content)
	}

	if qp.Language != "" {
		v.Set("language", qp.Language)
	}

	if qp.Status != "" {
		v.Set("status", string(qp.Status))
	}

	if qp.Category != "" {
		v.Set("category", qp.Category)
	}

	// Handle the Fields slice if it's provided
	if qp.Fields != nil && len(*qp.Fields) > 0 {
		// Convert each TemplateFields to its string representation
		fields := make([]string, len(*qp.Fields))
		for i, field := range *qp.Fields {
			fields[i] = string(field)
		}
		// Join the fields with commas
		v.Set("fields", strings.Join(fields, ","))
	}

	if qp.Summary != nil && len(*qp.Summary) > 0 {
		summary := make([]string, len(*qp.Summary))
		for i, s := range *qp.Summary {
			summary[i] = string(s)
		}
		v.Set("summary", strings.Join(summary, ","))
	}

	qp.GraphCursors.BuildQuery(&v)

	// Encode the query parameters and return the query string
	return v.Encode()
}
