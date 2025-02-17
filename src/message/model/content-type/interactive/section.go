package message_type_interactive_model

// Described at
//
// https://developers.facebook.com/docs/whatsapp/on-premises/reference/messages#section-object
type Section struct {
	Title        string         `json:"title,omitempty"`         // Required if the message has more than one section. Maximum length: 24 characters.
	Rows         *[]Row         `json:"rows,omitempty"`          // Required for List Messages. Contains a list of row objects. Limited to 10 rows across all sections.
	ProductItems *[]ProductItem `json:"product_items,omitempty"` // Required for Multi-Product Messages. There is a minimum of 1 product per section and a maximum of 30 products across all sections.
}
