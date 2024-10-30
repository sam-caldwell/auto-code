package arguments

// HelpText - Return the HelpText as a string
func (p *ParameterDefinition) HelpText() string {
	return string(p.helpText)
}
