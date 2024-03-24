package message_content_type_model

// Media messages.
type Media struct {
	Link string `json:"link,omitempty"` // To use link your media should be in a publicly available server.
	Id   string `json:"id,omitempty"`   // To use id your media should be in the meta server. You can upload an media to the meta server with https://developers.facebook.com/docs/whatsapp/cloud-api/reference/media#upload-media.
}
