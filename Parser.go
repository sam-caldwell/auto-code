package arguments

// Parser - A configurable multi-source configuration parser solution.
//
// The Parser object can be configured to parse various data sources and generate a Configuration object.
// The sources may include files, environment variables, command-line arguments or even a secrets manager.
//
// Parser must define the parameters of a program which will be compiled into the final configuration state.
// This information is stored in a map.
type Parser struct {
	programName        string
	programDescription string
	programCopyright   string
	parameters         ParameterCollection // Note: as a map, there is no guaranteed order.
	err                error
}
