package pdo

import (
	"github.com/sam-caldwell/auto-code/pkg/manifest"
	"github.com/sam-caldwell/auto-code/pkg/manifest/pdo/validator"
)

// DataObject - represents the actual dataContract value and its associated validation logic.
type DataObject struct {
	manifest.DataObjectWithReference
	State     any                 `yaml:"default,omitempty"`
	validator validator.Interface `yaml:"regex,omitempty"`
}
