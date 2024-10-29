package arguments

// HelpText - A type representing a parameter's help text string
type HelpText string

// String - Return the string version of a help text string
func (text *HelpText) String() string {
	return string(*text)
}

// set - A sanitizer/setter for the help text string.
func (text *HelpText) set(helpText *string) error {
	//ToDo: add validation
	*text = HelpText(*helpText)
	return nil
}
