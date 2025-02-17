package message_model

// When messages type is set to system, a customer has updated their phone number or profile information, this object is included in the messages object.
type System struct {
	Body     string     `json:"body"`      // Describes the change to the customer's identity or phone number.
	Identity string     `json:"identity"`  // Hash for the identity fetched from server.
	NewWaId  string     `json:"new_wa_id"` // New WhatsApp ID for the customer when their phone number is updated. Available on webhook versions v11.0 and earlier.
	WaId     string     `json:"wa_id"`     // New WhatsApp ID for the customer when their phone number is updated. Available on webhook versions v12.0 and later.
	Type     SystemType `json:"type"`      // Type of system update.
	Customer string     `json:"customer"`  // The WhatsApp ID for the customer prior to the update.
}

// Type of system update.
type SystemType struct {
	CustomerChangedNumber   string `json:"customer_changed_number "`  // A customer changed their phone number.
	CustomerIdentityChanged string `json:"customer_identity_changed"` // A customer changed their profile information.
}
