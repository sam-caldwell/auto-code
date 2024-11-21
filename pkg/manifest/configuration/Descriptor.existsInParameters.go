package configuration

// existsInParameters - returns whether the input element exists in the parameters list
func (c *Descriptor) existsInParameters(element string) bool {
	for _, parameter := range c.Parameters {
		if name := parameter.Name.String(); name == element {
			return true
		}
	}
	return false
}
