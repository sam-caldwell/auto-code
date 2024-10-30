package arguments

// String - Return the string version of a help text string
func (text *HelpText) String() string {
	return string(*text)
}
