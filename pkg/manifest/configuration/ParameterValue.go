package configuration

import (
	"github.com/sam-caldwell/auto-code/pkg/manifest/pdo"
	"github.com/sam-caldwell/auto-code/pkg/types/xref"
)

// ParameterValue - represents a parameter value, type and validation regex
//
// # This object supports external $ref references
//
// This object implements a DataObject (PDO) representing a strongly
// typed data value (state) and its validator logic.  While the PDO can store
// many different data types, it does perform type-checking and enforcement.
type ParameterValue struct {
	xref.DataObjectWithReference
	Data pdo.DataObject `yaml:"value"`
}
