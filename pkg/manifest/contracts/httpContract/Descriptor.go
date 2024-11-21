package httpContract

import (
	"github.com/sam-caldwell/auto-code/pkg/manifest/contracts/common"
	"github.com/sam-caldwell/auto-code/pkg/manifest/contracts/httpContract/httpDriver"
	"github.com/sam-caldwell/auto-code/pkg/types/xref"
)

// Descriptor - defines an HTTP contract
type Descriptor struct {
	xref.DataObjectWithReference
	common.Base
	Driver httpDriver.HttpDriver `yaml:"driver"`
	Schema []common.Schema       `yaml:"schema"`
}
