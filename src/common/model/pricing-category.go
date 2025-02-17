package common_model

type PricingCategory string

const (
	Authentication       PricingCategory = "authentication"        // Indicates the conversation was opened by a business sending template categorized as AUTHENTICATION to the customer. This applies any time it has been more than 24 hours since the last customer message.
	Marketing            PricingCategory = "marketing"             // Indicates the conversation was opened by a business sending template categorized as MARKETING to the customer. This applies any time it has been more than 24 hours since the last customer message.
	Utility              PricingCategory = "utility"               // Indicates the conversation was opened by a business sending template categorized as UTILITY to the customer. This applies any time it has been more than 24 hours since the last customer message.
	Service              PricingCategory = "service"               // Indicates that the conversation opened by a business replying to a customer within a customer service window.
	ReferralConversation PricingCategory = "referral_conversation" // Indicates a free entry point conversation.
)
