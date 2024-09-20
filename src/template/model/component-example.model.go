package template_model

type ComponentExample struct {
	HeaderText   *[]string   `json:"header_text,omitempty"`
	BodyText     *[][]string `json:"body_text,omitempty"`
	HeaderHandle *[]string   `json:"header_handle,omitempty"`
}
