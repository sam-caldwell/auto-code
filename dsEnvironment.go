package arguments

import (
	"fmt"
	"os"
	"regexp"
)

// dsEnvironment - A struct for representing an environment variable data source
//
// This struct should be created by the Environment() function.  It should NEVER be directly created.
type dsEnvironment struct {
	eName string
}

// Get - Given an environment variable name (eName) return its value.
//
// This method returns the unvalidated, raw value from an associated environment variable name (eName)
func (ds dsEnvironment) Get() (any, error) {
	value := os.Getenv(ds.eName)
	return value, nil
}

func Environment(eName string) DataSource {
	const pattern = `^[A-Za-z_][A-Za-z0-9_]*$`
	if regexp.MustCompile(pattern).MatchString(eName) {
		return dsEnvironment{
			eName: eName,
		}
	}
	return dsInvalidParameter{
		err: fmt.Errorf(invalidEnvironmentVariableName, eName),
	}
}
