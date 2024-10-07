package common_model

type Pricing struct {
	Category     PricingCategory `json:"category"`
	PricingModel string          `json:"pricing_model"`
	Billable     bool            `json:"billable"`
}
