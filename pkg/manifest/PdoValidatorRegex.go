package manifest

import "regexp"

// PdoValidatorRegex - a regular expression data object validator
type PdoValidatorRegex struct {
	re *regexp.Regexp
}
