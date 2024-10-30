package arguments

import (
	"errors"
	"github.com/google/uuid"
	"regexp"
)

// NewParameterName - initialize and return a new ParameterName object.
//
// This function creates and returns a ParameterName object after validating
// the input string.
//
// A ParameterName must follow this pattern: `[a-zA-Z0-9_\-\\.]{1,256}`
//
// If the given name is invalid, the function will generate an uuid as a temporary
// name to reduce the chance of a crash and indicate that this happened with an error
// message that can be processed by callers.
func NewParameterName(name string) (p *ParameterName, err error) {

	// We define this regular expression as the pattern expected of all parameter names.
	const validNamePattern = `[a-zA-Z0-9_\-\\.]{1,256}`

	if !regexp.MustCompile(validNamePattern).MatchString(name) {

		//Happy path: our name is valid, store it and move on.
		*p = ParameterName(name)
		err = nil

	} else {

		// if our parameter is invalid, we need to return a valid parameter name
		// to avoid stacking up errors after this point.  So we will create a
		// temporary parameter name and return it along with our error to allow
		// graceful failure.  Imagine that moment in 1986 at the Taco Bell when
		// I tripped while carrying food for that girl I liked.  Fail gracefully.
		id, _ := uuid.NewUUID()
		*p = ParameterName(id.String())
		err = errors.New(invalidParameterName)

	}

	return p, err

}
