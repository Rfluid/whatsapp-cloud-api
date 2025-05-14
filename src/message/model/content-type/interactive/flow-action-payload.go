package message_type_interactive_model

type FlowActionPayload struct {
	Screen string      `json:"screen,omitempty"` // Required. The id of the first screen of the Flow.
	Data   interface{} `json:"data,omitempty"`   // Optional. The input data for the first screen of the Flow. Must be a non-empty object.
}
