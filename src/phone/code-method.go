// Provides models to handle phone verifcation payload validation.
package phone

type CodeMethod string

const (
	SMS   CodeMethod = "SMS"
	Voice CodeMethod = "VOICE"
)
