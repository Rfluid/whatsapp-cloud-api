package message_type_template_model

type ComponentType string

const (
	Header  ComponentType = "HEADER"
	Body    ComponentType = "BODY"
	Footer  ComponentType = "FOOTER"
	Buttons ComponentType = "BUTTONS"
)
