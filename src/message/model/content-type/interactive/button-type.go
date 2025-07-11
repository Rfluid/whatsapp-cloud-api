package message_type_interactive_model

type ButtonType string

const Reply ButtonType = "reply"

// IsValid returns true if the ButtonType is one of the predefined constants.
func (bt ButtonType) IsValid() bool {
	switch bt {
	case Reply:
		return true
	default:
		return false
	}
}
