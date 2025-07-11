package template_model

type ComponentType string

const (
	Header  ComponentType = "HEADER"
	Body    ComponentType = "BODY"
	Footer  ComponentType = "FOOTER"
	Buttons ComponentType = "BUTTONS"
)

// IsValid checks if the ComponentType is a valid predefined value.
func (ct ComponentType) IsValid() bool {
	switch ct {
	case Header, Body, Footer, Buttons:
		return true
	default:
		return false
	}
}
