package business_model

// Business represents a business account with various attributes.
type Business struct {
	Id                       string `json:"id" query:"id"`                                                 // ID is the business account ID.
	Name                     string `json:"name" query:"name"`                                             // Name is the name of the business.
	Currency                 string `json:"currency" query:"currency"`                                     // Currency is the currency used by the business.
	TimezoneId               string `json:"timezone_id" query:"timezone_id"`                               // TimezoneID represents the timezone of the business.
	MessageTemplateNamespace string `json:"message_template_namespace" query:"message_template_namespace"` // MessageTemplateNamespace represents the namespace of the message template.
}
