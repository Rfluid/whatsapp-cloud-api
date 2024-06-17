package message_model

type (
	ReceiveType string
	Type        ReceiveType
)

const (
	Text        Type = "text"
	Reaction    Type = "reaction"
	Location    Type = "location"
	Contacts    Type = "contacts"
	Interactive Type = "interactive"
	Template    Type = "template"

	Image    Type = "image"
	Video    Type = "video"
	Audio    Type = "audio"
	Sticker  Type = "sticker"
	Document Type = "document"

	Button Type = "button"
	Order  Type = "order"
)

const (
	SystemReceiveType ReceiveType = "system"
	Unknown           ReceiveType = "unknown"
)
