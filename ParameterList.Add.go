package arguments

import "fmt"

// Add - Append a parameter (by reference) to the parameter map.
func (p *ParameterList) Add(parameter *ParameterDefinition) error {
	if _, ok := (*p)[parameter]; ok {
		return fmt.Errorf("duplicate parameter (%s)", *parameter.name)
	}
	*p[parameter] = struct{}{}
	return nil
}

//ToDo: verify this is even needed
