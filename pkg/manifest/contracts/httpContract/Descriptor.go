package httpContract

import (
	"github.com/sam-caldwell/auto-code/pkg/manifest/contracts/common"
	"github.com/sam-caldwell/auto-code/pkg/types/generic/enum"
	"github.com/sam-caldwell/auto-code/pkg/types/xref"
)

// Descriptor - defines an HTTP contract
type Descriptor struct {
	xref.DataObjectWithReference
	common.Base
	Driver enum.HttpDriver `yaml:"driver"`
	Schema []common.Schema `yaml:"schema"`
}
