package arguments

import "fmt"

func (p *ParameterCollection) add(parameterName *ParameterName, helpText *HelpText,
	parameterDefinition *ParameterDefinition) (err error) {

	if parameterName == nil {
		// Create a default (uuid) parameter name and ignore the expected error in doing this.
		parameterName, _ = NewParameterName("")
		return fmt.Errorf(nilParameterName)
	}
	if parameterDefinition == nil {
		// create a default ParameterDefinition and ignore the expected error.
		parameterDefinition, _ = NewParameterDefinition(nil, nil, nil)
		return fmt.Errorf(nilParameterDefinition)
	}

	(*p)[*parameterName] = parameterDefinition

}
