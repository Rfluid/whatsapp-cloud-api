package message_type_common_model

// Location messages.
type LocationData struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Name      string  `json:"name"`
	Address   string  `json:"address,omitempty"`
}
