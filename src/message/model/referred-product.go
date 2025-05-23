package message_model

// Referred product object describing the product the user is requesting information about
type ReferredProduct struct {
	CatalogId         string `json:"catalog_id"`          // Unique identifier of the Meta catalog linked to the WhatsApp Business Account.
	ProductRetailerId string `json:"product_retailer_id"` // Unique identifier of the product in a catalog.
}
