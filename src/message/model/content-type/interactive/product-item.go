package message_type_interactive_model

type ProductItem struct {
	ProductRetailerId string `json:"product_retailer_id"` // Required for Multi-Product Messages. Unique identifier of the product in a catalog. To get this ID, go to Commerce Manager, select your account and the shop you want to use. Then, click Catalog > Items, and find the item you want to mention. The ID for that item is displayed under the item's name.
}
