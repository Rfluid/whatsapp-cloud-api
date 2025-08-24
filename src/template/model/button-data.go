package template_model

type ButtonData struct {
	Type ButtonType `json:"type" validate:"required,template_button_type"`

	Text        string `json:"text,omitempty"`
	Url         string `json:"url,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Example     any    `json:"example,omitempty"` // Check valid examples here https://developers.facebook.com/docs/whatsapp/business-management-api/message-templates/#template-components

	FlowID         int    `json:"flow_id,omitempty"`
	FlowAction     string `json:"flow_action,omitempty"`
	NavigateScreen string `json:"navigate_screen,omitempty"`
}
