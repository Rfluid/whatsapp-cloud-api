package media_model

// Media messages.
type UseMedia struct {
	Link     string `json:"link,omitempty"`     // To use link your media should be in a publicly available server.
	ID       string `json:"id,omitempty"`       // To use id your media should be in the meta server. You can upload an media to the meta server with https://developers.facebook.com/docs/whatsapp/cloud-api/reference/media#upload-media.
	Caption  string `json:"caption,omitempty"`  // Optional and not used for audio or sticker media.
	Filename string `json:"filename,omitempty"` // Optional and used for document media.
}
