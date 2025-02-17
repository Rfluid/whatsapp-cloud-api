package template_model

type ButtonData struct {
	Type ButtonType `json:"type"`

	Text        string      `json:"text,omitempty"`
	Url         string      `json:"url,omitempty"`
	PhoneNumber string      `json:"phone_number,omitempty"`
	Example     interface{} `json:"example,omitempty"` // string or string array

	FlowId         string `json:"flow_id,omitempty"`
	FlowAction     string `json:"flow_action,omitempty"`
	NavigateScreen string `json:"navigate_screen,omitempty"`
}
