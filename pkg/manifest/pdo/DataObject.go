package pdo

import (
	"github.com/sam-caldwell/auto-code/pkg/manifest/pdo/validator"
	"github.com/sam-caldwell/auto-code/pkg/types/xref"
)

// DataObject - represents the actual dataContract value and its associated validation logic.
type DataObject struct {
	xref.DataObjectWithReference
	State     any                 `yaml:"default,omitempty"`
	validator validator.Interface `yaml:"regex,omitempty"`
}
