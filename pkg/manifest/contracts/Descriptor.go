package contracts

import (
	"github.com/sam-caldwell/auto-code/pkg/manifest/contracts/dataContract"
	"github.com/sam-caldwell/auto-code/pkg/manifest/contracts/httpContract"
)

// Descriptor - represents the contract types
type Descriptor struct {
	Data dataContract.Descriptor `yaml:"data"`
	Http httpContract.Descriptor `yaml:"http"`
}
