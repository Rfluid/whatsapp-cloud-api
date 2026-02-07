package common

type Error struct {
	Message      string `json:"message"` // A combination of the error code and title.
	Type         string `json:"type"`    // Type of error. Example: OAuthException
	Code         int    `json:"code"`    // Error codes (not HTPP codes) described at https://developers.facebook.com/docs/whatsapp/cloud-api/support/error-codes/
	ErrorSubcode *int   `json:"error_subcode,omitempty"`

	ErrorData *ErrorData `json:"error_data,omitempty"`

	IsTransient    *bool  `json:"is_transient,omitempty"`
	ErrorUserTitle string `json:"error_user_title,omitempty"`
	ErrorUserMsg   string `json:"error_user_msg,omitempty"`
	FBTraceID      string `json:"fbtrace_id"` // Unique identifier for the error. Use this ID when contacting support.
}

type ErrorData struct {
	Details string `json:"details"` // Describes the error and provides most probable reason. Might contain information about how to solve the error. Example: Message failed to send because there were too many messages sent from this phone number in a short period of time.
	MessagingProduct
}

type ErrorResponse struct {
	ErrorField Error `json:"error"`
}

func (e *Error) Error() string {
	return e.Message
}

func (e *ErrorResponse) Error() string {
	return e.ErrorField.Message
}
