package validator

// Interface - is used to allow omitempty behavior for different types of parameters.
//
// This is implemented by validator.Pattern, validator.Int* and validator.Uint* as well as validator.Null
type Interface interface {
	Verify() error
}
