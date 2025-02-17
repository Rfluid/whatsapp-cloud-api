package message_type_common_model

type ProductItem struct {
	ProductRetailerId string `json:"product_retailer_id"`
	Quantity          string `json:"quantity"`
	ItemPrice         string `json:"item_price"`
	Currency          string `json:"currency"`
}
