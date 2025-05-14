package phone_verification_model

import "net/url"

type RequestCode struct {
	CodeMethod CodeMethod `json:"code_method"` // Verification method.
	Language   string     `json:"language"`    // Two characters language code.
}

func (r *RequestCode) ToURLValues() url.Values {
	formData := url.Values{}

	formData.Set("code_method", string(r.CodeMethod))
	formData.Set("language", r.Language)

	return formData
}
