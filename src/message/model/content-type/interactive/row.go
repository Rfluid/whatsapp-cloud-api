package message_type_interactive_model

type Row struct {
	Title       string `json:"title"`                 // Required. Maximum length: 24 characters.
	ID          string `json:"id"`                    // Required. Maximum length: 200 characters.
	Description string `json:"description,omitempty"` // Optional. Maximum length: 72 characters.
}
