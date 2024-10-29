package arguments

import "fmt"

// Configuration - An object representing the final configuration state of an application.
//
// The Configuration object is the final read-only data object containing the resolved configuration state of the
// program based on its parser definition's interpretation.  The Configuration is created by the Parser object's
// Parse() method.
type Configuration struct {
	data map[string]any
}

// set - set configuration key-value pair (unless reserved word is detected)
//
// This method is not exported intentionally as the Configuration should be read-only outside this package.
func (config *Configuration) set(name string, value any) error {
	config.data[name] = value
	return nil
}

// Get - an exported function used to read data from the configuration object
func (config *Configuration) Get(key string) (any, error) {
	if value, ok := config.data[key]; ok {
		return value, nil
	}
	return nil, fmt.Errorf("not found")
}
