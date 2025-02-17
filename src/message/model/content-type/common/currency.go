package message_type_common_model

type Currency struct {
	FallbackValue string `json:"fallback_value"` // Text in case it fails.
	Code          string `json:"code"`           // Use ISO 4217.
	Amount1000    int    `json:"amount_1000"`    // Value times 1000.
}
