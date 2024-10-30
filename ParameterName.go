package arguments

// ParameterName - the name of a given configuration parameter
//
// Please use NewParameterName() to create ParameterName objects to ensure
// proper validation and avoid otherwise avoidable bugs.
//
// A ParameterName must follow the pattern enforced by NewParameterName()
type ParameterName string
