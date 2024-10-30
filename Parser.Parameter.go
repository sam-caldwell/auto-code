package arguments

// Parameter - Create a new Parameter definition and return a Parser object (by reference).
//
// A Parser Parameter is an object which has a name, help text, Value object and a set
// of one or more ParameterDefinition objects.  This defines the configuration parameter
// as a term (name), data value (Value) and associated validator logic as well as a set
// of data sources (DataSourceCollection), which are applied in the order in which
// they appear in the Parameter() method call.
//
// The Value object represents the default value (and final value after Parser.Parse() is called)
// as well as the value's validator logic. If *Value is nil, then no default value is defined and
// no validation logic exists.  Yolo!
//
// The DataSourceCollection describes various supported configuration inputs, such as YamlFile,
// JsonFile, KeyValueFile, Environment (variable), Cli (command-line args) or supported secrets
// managers.  This collection is designed to be extensible as additional data sources are added.
//
// Example:
//
//	config, err := arguments.NewParser().
//		Parameter(
//			"network.address",                                              // name the parameter
//			"this is the listener Ip address",                              // give the parameter some help text
//			arguments.String("0.0.0.0", arguments.Pattern(ipAddressRegex)), // add a validator
//			arguments.YamlFile("config.yaml", "net.ip"),                    // add at least one supported data source
//			arguments.Environment("LISTEN_IP"),
//			arguments.Cli("-n"),
//			arguments.Cli("--ip"),
//		).
//		Parse()                              // process the definition to create a fully resolved Configuration.
func (parser *Parser) Parameter(name, help string, value ValueDataInterface, sources ...DataSource) *Parser {

	// initialize our data sources, which may cause an error we need to register.
	dataSources, err := NewDataSourceCollection(&sources)

	if err != nil {
		// We register the error state, but we do not bail,
		// since that could create other errors or a crash.
		// This may mean we are passing an empty (non-nil)
		// DataSourceCollection to NewParameterCollection().
		parser.err = err
	}

	// initialize our parameter collection, which may cause an error we need to register.
	parser.parameters, err = NewParameterCollection(name, help, value, dataSources)

	if err != nil && parser.err == nil {
		// We don't want to overwrite the first error since it probably is causing
		// the second error.  Instead, we will ignore this error and the user will
		// encounter this one once the first error is resolved.
		parser.err = err
	}

	return parser

}
