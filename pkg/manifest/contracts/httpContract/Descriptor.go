package httpContract

import (
	"github.com/sam-caldwell/auto-code/pkg/manifest/contracts/common"
)

// Descriptor - defines an HTTP contract
type Descriptor struct {
	common.Base
	Schema []Schema `yaml:"endpoints"`
}
