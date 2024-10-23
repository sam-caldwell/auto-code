package arguments

const (
	validationFailedError                         = "validation failed.  Must have follow '%s'"
	numericTypeCannotPerformPatternBoundsCheck    = "numeric type cannot perform pattern bounds check"
	nonNumericTypeCannotPerformNumericBoundsCheck = "non-numeric type cannot perform pattern bounds check"
	nilValuePassedToNumericBoundary               = "nil value passed to NumericBoundary"
	boundsCheckFailed                             = "bounds check for %s failed (%v)"
	defaultValueMustPassValidation                = "default value must pass validation (%s)"
)

const (
	NoValidation = ``
)
