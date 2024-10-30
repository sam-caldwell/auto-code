package arguments

// NewParameterCollection - initializes and returns a new ParameterCollection.
func NewParameterCollection(name, help string, value *Value, sources []*DataSource) (ParameterCollection, error) {
	pName, err := NewParameterName(name)
	pHelp, err := NewHelpText

	p := make(ParameterCollection)
	p.add(name, helpText, value, sources)
	return p, nil
}
