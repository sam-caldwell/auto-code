package arguments

const (
	emptyDataSourceList            = "parameter definitions require at least one DataSource"
	invalidCliOption               = "a cli option must start with - or -- and have at least one character or number (%s)"
	invalidDefaultValue            = "invalid default value"
	invalidEnvironmentVariableName = "an invalid environment variable name is encountered (%s)"
	invalidHelpText                = "invalid help text (cannot be empty or only whitespace"
	invalidParameterName           = "invalid parameter name"
	invalidProgramName             = "invalid program name"
	invalidValueType               = "unknown or unexpected value type"
	nilParameterDefinition         = "parameter definition is nil"
	nilParameterName               = "parameterName is nil"
	nilDataSourceListPointer       = "programming error: nil sources pointer"
	nilSourceEncountered           = "a nil data source was encountered in arguments parser"
	nilValidatorObject             = "a nil validator object was encountered"
	nilValueObject                 = "a nil value object was encountered in arguments parser"

	validationFailedError                         = "validation failed.  Must have follow '%s'"
	numericTypeCannotPerformPatternBoundsCheck    = "numeric type cannot perform pattern bounds check"
	nonNumericTypeCannotPerformNumericBoundsCheck = "non-numeric type cannot perform pattern bounds check"
	nilValuePassedToNumericBoundary               = "nil value passed to NumericBoundary"
	boundsCheckFailed                             = "bounds check for %s failed (%v)"
	defaultValueMustPassValidation                = "default value must pass validation (%s)"
	unknownOrUnexpectedArgument                   = "unknown or unexpected argument"
	unsupportedArgumentClass                      = "unsupported argument class"
	invalidBooleanValue                           = "invalid boolean value (expect true|false)"
)

const (
	NoArgumentExpected = ""
	NoValidation       = ``
)
