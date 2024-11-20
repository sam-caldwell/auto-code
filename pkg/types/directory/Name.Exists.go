package directory

import (
	parent "github.com/sam-caldwell/directory"
)

// Exists - returns boolean true if the current directory exists
func (p *Name) Exists() bool {
	return parent.Exists(string(*p))
}
