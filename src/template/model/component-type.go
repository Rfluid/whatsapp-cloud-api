package template_model

type ComponentType string

const (
	Header          ComponentType = "HEADER"
	Body            ComponentType = "BODY"
	Footer          ComponentType = "FOOTER"
	ButtonComponent ComponentType = "BUTTON"
)

// IsValid checks if the ComponentType is a valid predefined value.
func (ct ComponentType) IsValid() bool {
	switch ct {
	case Header, Body, Footer, ButtonComponent:
		return true
	default:
		return false
	}
}
