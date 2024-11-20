package directory

import (
	parent "github.com/sam-caldwell/directory"
)

// Exists - returns boolean true if the given directory exists
func Exists(p string) bool {
	return parent.Exists(p)
}
