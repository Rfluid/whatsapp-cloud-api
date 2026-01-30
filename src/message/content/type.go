package content

type (
	ReceiveType string
	Type        ReceiveType
)

const (
	Text        Type = "text"
	Reaction    Type = "reaction"
	Location    Type = "location"
	Contacts    Type = "contacts"
	InteractiveMsg Type = "interactive"
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

// IsValid checks if the Type is one of the supported message types.
func (t Type) IsValid() bool {
	switch t {
	case Text, Reaction, Location, Contacts, InteractiveMsg, Template,
		Image, Video, Audio, Sticker, Document,
		Button, Order:
		return true
	default:
		return false
	}
}
