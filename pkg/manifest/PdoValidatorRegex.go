package manifest

import "regexp"

// PdoValidatorRegex - a regular expression dataContract object validator
type PdoValidatorRegex struct {
	re *regexp.Regexp
}
