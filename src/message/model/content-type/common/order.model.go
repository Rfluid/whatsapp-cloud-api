package message_type_common_model

// Included in the messages object when a customer has placed an order.
type Order struct {
	CatalogId    string        `json:"catalog_id"` // ID for the catalog the ordered item belongs to.
	Text         string        `json:"text"`       // Text message from the user sent along with the order.
	ProductItems []ProductItem `json:"product_items"`
}
