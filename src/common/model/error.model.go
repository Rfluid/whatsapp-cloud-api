package common_model

type Error struct {
	Message   string    `json:"message"` // A combination of the error code and title.
	Type      string    `json:"type"`    // Type of error. Example: OAuthException
	Code      int       `json:"code"`    // Error codes (not HTPP codes) described at https://developers.facebook.com/docs/whatsapp/cloud-api/support/error-codes/
	ErrorData ErrorData `json:"error_data"`
	FBTraceID string    `json:"fbtrace_id"` // Unique identifier for the error. Use this ID when contacting support.
}

type ErrorData struct {
	Details string `json:"details"` // Describes the error and provides most probable reason. Might contain information about how to solve the error. Example: Message failed to send because there were too many messages sent from this phone number in a short period of time.
	MessagingProduct
}

func (e *Error) Error() string {
	return e.Message
}
