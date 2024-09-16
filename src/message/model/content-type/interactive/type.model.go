package message_type_interactive_model

type InteractiveType string

// Described at
//
// https://developers.facebook.com/docs/whatsapp/on-premises/reference/messages#interactive-object
var (
	List           InteractiveType = "list"
	Button         InteractiveType = "button"
	Product        InteractiveType = "product"
	ProductList    InteractiveType = "product_list"
	CatalogMessage InteractiveType = "catalog_message"
	Flow           InteractiveType = "flow"
)
