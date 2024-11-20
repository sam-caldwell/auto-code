package file

import (
	parent "github.com/sam-caldwell/file"
)

// Exists - returns boolean true if the current file exists
func (f *Name) Exists() bool {
	return parent.Exists(string(*f))
}
