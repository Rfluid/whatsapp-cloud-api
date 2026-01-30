package phone

import "net/url"

type VerifyCodePayload struct {
	Code string `json:"code"`
}

func (v *VerifyCodePayload) ToURLValues() url.Values {
	formData := url.Values{}

	formData.Set("code", v.Code)

	return formData
}
