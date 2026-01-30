package interactive

type ButtonReplyData struct {
	ID    string `json:"id,omitempty"`    // A unique identifier for the button reply.
	Title string `json:"title,omitempty"` // The title of the button reply.
}
