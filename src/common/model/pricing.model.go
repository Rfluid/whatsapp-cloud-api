package common_model

type Pricing struct {
	Category     *PricingCategory `json:"category,omitempty"`
	PricingModel *string          `json:"pricing_model,omitempty"`
	Billable     *bool            `json:"billable,omitempty"`
}
