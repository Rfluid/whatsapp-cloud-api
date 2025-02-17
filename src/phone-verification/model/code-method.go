// Provides models to handle phone verifcation payload validation.
package phone_verification_model

type CodeMethod string

const (
	SMS   CodeMethod = "SMS"
	Voice CodeMethod = "VOICE"
)
