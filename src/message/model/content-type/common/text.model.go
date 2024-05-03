package message_type_common_model

type Text struct {
	PreviewURL bool   `json:"preview_url"`
	Body       string `json:"body"`
}
