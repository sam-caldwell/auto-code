package arguments

import "fmt"

// NewParameterDefinition - Create a new ParameterDefinition object with proper validation.
//
// A new ParameterDefinition object should make sure sources have an initialized sources map
// and otherwise validate the data.
//
// This function assumes helpText and sources were created using their proper initializer functions,
// but we still perform nil pointer checks to ensure that changes have not left us vulnerable to
// a bad state that could crash the system. If the nil pointer situation occurs, the bad state will
// be "fixed" temporarily with a default state object and the error will be reported.
func NewParameterDefinition(helpText *HelpText, value ValueDataInterface,
	sources DataSourceCollection) (p *ParameterDefinition, err error) {

	if nil == helpText {
		helpText = new(HelpText)
		err = fmt.Errorf(invalidHelpText)
	}

	if sources == nil {
		sources = make(DataSourceCollection, 0)
		err = fmt.Errorf(nilDataSourceListPointer)
	}

	p = &ParameterDefinition{
		helpText: *helpText,
		value:    value,
		sources:  sources,
	}

	return p, err
}
