package common_model

type Error struct {
	Code      int       `json:"code"`
	Title     string    `json:"title"`
	Message   string    `json:"message"` // This value is the same as the title value. For example: Rate limit hit. Note that the message property in API error response payloads pre-pends this value with the a # symbol and the error code in parenthesis. For example: (#130429) Rate limit hit.
	ErrorData ErrorData `json:"error_data"`
}

type ErrorData struct {
	Details string `json:"details"` // Describes the error. Example: Message failed to send because there were too many messages sent from this phone number in a short period of time.
	any
}
