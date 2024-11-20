package pdo

import "github.com/sam-caldwell/auto-code/pkg/manifest"

// ParameterDataObject - represents the actual dataContract value and its associated validation logic.
type ParameterDataObject struct {
	manifest.DataObjectWithReference
	State     any              `yaml:"default,omitempty"`
	validator PdoDataValidator `yaml:"regex,omitempty"`
}
