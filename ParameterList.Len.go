package arguments

// Len - Return the length of the parameter map
func (p *ParameterList) Len() int {
	return len(*p)
}

//ToDo: verify this is even needed
