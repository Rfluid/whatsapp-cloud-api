package message_type_contact_model

type Phone struct {
	Phone string `json:"phone"`
	Type  string `json:"type"`
	WaID  string `json:"wa_id,omitempty"`
}
