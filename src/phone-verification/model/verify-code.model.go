package phone_verification_model

import "net/url"

type VerifyCode struct {
	Code string `json:"code"`
}

func (v *VerifyCode) ToURLValues() url.Values {
	formData := url.Values{}

	formData.Set("code", v.Code)

	return formData
}
