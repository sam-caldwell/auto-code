package arguments

// ParameterName - the name of a given configuration parameter
type ParameterName string

// set - a sanitizer/setter for the ParameterName type
func (p *ParameterName) set(name string) error {
	//ToDo: add validation and return errors if appropriate
	*p = ParameterName(name)
	return nil
}
