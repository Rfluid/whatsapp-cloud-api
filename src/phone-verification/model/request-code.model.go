package phone_verification_model

type RequestCode struct {
	CodeMethod CodeMethod `json:"code_method"` // Verification method.
	Language   string     `json:"language"`    // Two characters language code.
}
