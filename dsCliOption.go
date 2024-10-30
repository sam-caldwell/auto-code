package arguments

// dsCliOption - Defines the data source from which a command-line argument can be extracted.
type dsCliOption struct{}

// Get - Returns the data object's value and any error state
func (ds dsCliOption) Get() (any, error) {
	return nil, nil
}

// CliOption - Create a command-line option data source definition.
//
// This function creates and configures a dsCliOption object, which will be used to parse a command-line option
// and expose the validated value for use by the Parser.Parse() method.
func CliOption(name, help string, validator *Validator) DataSource {
	return dsCliOption{}
}
