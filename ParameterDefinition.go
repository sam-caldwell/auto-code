package arguments

// ParameterDefinition - a configuration parameter definition.
//
//	The ParameterDefinition object defines the parameter, providing a helpful HelpText object, a Value object pointer
//	and a DataSourceCollection from which the parameter's actual value can be obtained.
//
//	A ParameterDefinition assumes that--
//	   - HelpText was created with NewHelpText(),
//	   - Value (pointer) is created with NewValue()
//	   - DataSourceCollection is created with NewDataSourceCollection()
//	since these functions perform data validation.
//
//	ParameterDefinition
//	  │
//	  ├── helpText:  HelpText
//	  ├── value: *Value
//	  └── sources: DataSourceCollection
type ParameterDefinition struct {
	helpText HelpText
	value    *Value
	sources  DataSourceCollection // order must be preserved!
}
