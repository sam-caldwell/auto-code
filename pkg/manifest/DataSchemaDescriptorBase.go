package manifest

import (
	"github.com/sam-caldwell/auto-code/pkg/manifest/dataDriver"
	"github.com/sam-caldwell/auto-code/pkg/manifest/dataDriver/dataCommon"
)

// DataSchemaDescriptorBase - a base class used to create common functionality
//
// This is a 'base class' per se to define common functionality.
type DataSchemaDescriptorBase struct {
	Driver dataDriver.DataDriver `yaml:"driver"`

	ToolChain []dataCommon.ToolChainCodeBlocks
}
