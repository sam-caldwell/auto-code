package dataContract

import (
	"github.com/sam-caldwell/auto-code/pkg/manifest/contracts/common"
	"github.com/sam-caldwell/auto-code/pkg/types/generic/enum"
	"github.com/sam-caldwell/auto-code/pkg/types/xref"
)

// Descriptor - represents the content of a dataContract contract
type Descriptor struct {
	xref.DataObjectWithReference
	common.Base
	Driver enum.DataDriver `yaml:"driver"`
	Schema common.Schema   `yaml:"schema"`
}
