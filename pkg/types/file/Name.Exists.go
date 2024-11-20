package file

import (
	parent "github.com/sam-caldwell/file"
)

// Exists - returns boolean true if the given file exists
func Exists(p string) bool {
	return parent.Exists(p)
}
