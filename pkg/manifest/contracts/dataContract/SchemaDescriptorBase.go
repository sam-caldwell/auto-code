package dataContract

import (
	"github.com/sam-caldwell/auto-code/pkg/manifest/dataDriver"
	"github.com/sam-caldwell/auto-code/pkg/manifest/dataDriver/dataCommon"
)

// SchemaDescriptorBase - a base class used to create common functionality
//
// This is a 'base class' per se to define common functionality.
type SchemaDescriptorBase struct {
	Driver dataDriver.DataDriver `yaml:"driver"`

	ToolChain []dataCommon.ToolChainCodeBlocks
}
