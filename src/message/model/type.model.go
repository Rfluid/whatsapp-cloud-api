package message_model

type Type string

const (
	Text        Type = "text"
	Reaction    Type = "reaction"
	Location    Type = "location"
	Contacts    Type = "contacts"
	Interactive Type = "interactive"

	Image    Type = "image"
	Video    Type = "video"
	Audio    Type = "audio"
	Sticker  Type = "sticker"
	Document Type = "document"
)
