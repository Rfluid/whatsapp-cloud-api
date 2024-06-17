package webhook_model

// Contact object with information for the customer who sent a message to the business.
type Contact struct {
	WaId    string `json:"wa_id"`
	Profile string `json:"profile"`
}
