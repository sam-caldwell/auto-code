package pdo

import "github.com/sam-caldwell/auto-code/pkg/manifest"

// DataObject - represents the actual dataContract value and its associated validation logic.
type DataObject struct {
	manifest.DataObjectWithReference
	State     any       `yaml:"default,omitempty"`
	validator Validator `yaml:"regex,omitempty"`
}
