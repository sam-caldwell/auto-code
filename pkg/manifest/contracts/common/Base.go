package common

import (
	"github.com/sam-caldwell/auto-code/pkg/types/generic"
	"github.com/sam-caldwell/auto-code/pkg/types/xref"
)

// Base - represents common contract elements
type Base struct {
	xref.DataObjectWithReference
	Name      generic.Identifier   `yaml:"name"`
	Artifacts []generic.Identifier `yaml:"artifacts"`
}
