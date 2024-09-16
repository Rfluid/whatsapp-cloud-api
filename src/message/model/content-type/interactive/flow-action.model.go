package message_type_interactive_model

// Use navigate to predefine the first screen as part of the message. Use data_exchange for advanced use-cases where the first screen is provided by your endpoint.
type FlowAction string

var (
	Navigate     FlowAction = "navigate"
	DataExchange FlowAction = "data_exchange"
)
