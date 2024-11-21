package validator

import "regexp"

// Regex - a regular expression dataContract object Interface
type Regex struct {
	re *regexp.Regexp
}
