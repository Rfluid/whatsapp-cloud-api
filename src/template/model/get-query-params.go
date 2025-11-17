package template_model

import (
	"net/url"
	"strconv"

	common_model "github.com/Rfluid/whatsapp-cloud-api/src/common/model"
)

type TemplateQueryParams struct {
	Category      TemplateCategory     `json:"category,omitempty" validate:"omitempty,template_category"`
	Content       string               `json:"content,omitempty"`
	Language      string               `json:"language,omitempty"`
	Name          string               `json:"name,omitempty"`
	NameOrContent string               `json:"name_or_content,omitempty" query:"name_or_content"`
	QualityScore  TemplateQualityScore `json:"quality_score,omitempty" query:"quality_score" validate:"omitempty,template_quality_score"`
	Status        Status               `json:"status,omitempty" validate:"omitempty,template_status"`

	Limit *uint64 `json:"limit,omitempty" query:"limit"`

	common_model.GraphCursors
}

// BuildQuery constructs the query string from TemplateQueryParams.
// It returns an encoded query string that can be appended to a URL.
func (qp *TemplateQueryParams) BuildQuery() string {
	// Initialize a URL values object
	v := url.Values{}

	// Conditionally set each query parameter if it's provided

	// Template fields
	if qp.Category != "" {
		v.Set("category", string(qp.Category))
	}

	if qp.Content != "" {
		v.Set("components", qp.Content)
	}

	if qp.Language != "" {
		v.Set("language", qp.Language)
	}

	if qp.Name != "" {
		v.Set("name", qp.Name)
	}

	if qp.NameOrContent != "" {
		v.Set("name_or_content", qp.Name)
	}

	if qp.QualityScore != "" {
		v.Set("quality_score", string(qp.QualityScore))
	}

	if qp.Status != "" {
		v.Set("status", string(qp.Status))
	}

	// Pagination fields
	if qp.Limit != nil {
		v.Set("limit", strconv.FormatUint(*qp.Limit, 10))
	}

	// Graph cursor fields
	qp.GraphCursors.BuildQuery(&v)

	// Encode the query parameters and return the query string
	return v.Encode()
}
