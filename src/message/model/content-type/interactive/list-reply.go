package message_type_interactive_model

type ListReplyData struct {
	Id          string `json:"id,omitempty"`          // A unique identifier for the list reply.
	Title       string `json:"title"`                 // The title of the list reply.
	Description string `json:"description,omitempty"` // The description of the list reply.
}
