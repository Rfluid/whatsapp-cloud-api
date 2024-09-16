package message_type_interactive_model

// Described at
//
// https://developers.facebook.com/docs/whatsapp/on-premises/reference/messages#action-object
type Action struct {
	Button            string        `json:"button,omitempty"`              // Required for list messages. It cannot be an empty string and must be unique within the message. Emojis are supported, markdown is not. Maximum length: 20 characters.
	Buttons           *[]ButtonData `json:"buttons,omitempty"`             // Required for Reply Button Messages.
	Sections          *[]Section    `json:"sections,omitempty"`            // Required for List Messages and Multi-Product Messages. There is a minimum of 1 and maximum of 10.
	CatalogId         string        `json:"catalog_id,omitempty"`          // Required for Single-Product Messages and Multi-Product Messages. Unique identifier of the Facebook catalog linked to your WhatsApp Business Account. This ID can be retrieved via Commerce Manager.
	ProductRetailerId string        `json:"product_retailer_id,omitempty"` // Required for Single-Product Messages and Multi-Product Messages. Unique identifier of the product in a catalog. Maximum 100 characters for both Single-Product and Multi-Product messages. To get this ID, go to Commerce Manager, select your Facebook Business account, and you will see a list of shops connected to your account. Click the shop you want to use. On the left-side panel, click Catalog > Items, and find the item you want to mention. The ID for that item is displayed under the item's name.

	*FlowData
}
