package arguments

// NewParameterCollection - initializes and returns a new ParameterCollection.
func NewParameterCollection(name, help string, value ValueDataInterface,
	sources []*DataSource) (p ParameterCollection, err error) {

	p = make(ParameterCollection)

	pName, err := NewParameterName(name)
	if err != nil {
		return p, err
	}
	pHelp, err := NewHelpText(&help)
	if err != nil {
		return p, err
	}

	parameterDefinition, err := NewParameterDefinition(pHelp, value, sources)
	if err != nil {
		return p, err
	}

	err = p.add(pName, pHelp, parameterDefinition)
	return p, err
}
