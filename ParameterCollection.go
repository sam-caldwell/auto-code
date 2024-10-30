package arguments

// ParameterCollection - a map of Parameters defined for the arguments.Parser
//
// The ParameterCollection type is a map of a ParameterName to a pointer to a ParameterDefinition,
// where ParameterName is the name used by the application to access a given parameter's Parser
// definition or final Configuration value.
//
// Note that we use the ParameterDefinition pointer in order to update the map easier.
//
// This object should be initialized with NewParameterCollection() for safety.
type ParameterCollection map[ParameterName]*ParameterDefinition
