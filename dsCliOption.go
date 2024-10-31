package arguments

import (
	"fmt"
	"regexp"
)

// dsCliOption - Defines the data source from which a command-line argument can be extracted.
type dsCliOption struct {
}

// Get - Returns the data object's value and any error state
func (ds dsCliOption) Get() (any, error) {
	return nil, nil
}

// CliOption - Create a command-line option data source definition.
//
// This function creates and configures a dsCliOption object, which will be used to parse a command-line option
// and expose the validated value for use by the Parser.Parse() method.
func CliOption(name string) DataSource {
	const pattern = `^-{1,2}[a-zA-Z0-9]{1,64}$`
	if regexp.MustCompile(pattern).MatchString(name) {
		return dsCliOption{}
	} else {
		// Return a nil, which will cause an error downstream when the
		//
		return dsInvalidParameter{
			err: fmt.Errorf(invalidCliOption, name),
		}
	}
}
