package phone

import "net/url"

type RequestCodePayload struct {
	CodeMethod CodeMethod `json:"code_method"` // Verification method.
	Language   string     `json:"language"`    // Two characters language code.
}

func (r *RequestCodePayload) ToURLValues() url.Values {
	formData := url.Values{}

	formData.Set("code_method", string(r.CodeMethod))
	formData.Set("language", r.Language)

	return formData
}
