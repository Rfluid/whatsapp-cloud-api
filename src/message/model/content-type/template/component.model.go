package message_type_template_model

type Component struct {
	Type       ComponentType `json:"type"`
	Parameters []Parameter   `json:"parameters"`
}
