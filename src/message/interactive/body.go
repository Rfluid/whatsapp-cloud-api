package interactive

// Described at
//
// https://developers.facebook.com/docs/whatsapp/on-premises/reference/messages#body-object
type Body struct {
	Text *string `json:"text,omitempty"` // Required if body is present
}
