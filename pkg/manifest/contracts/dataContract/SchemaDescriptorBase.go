package dataContract

import (
	"github.com/sam-caldwell/auto-code/pkg/manifest/contracts/dataContract/dataDriver"
	"github.com/sam-caldwell/auto-code/pkg/manifest/contracts/dataContract/dataDriver/dataCommon"
)

// SchemaDescriptorBase - a base class used to create common functionality
//
// This is a 'base class' per se to define common functionality.
type SchemaDescriptorBase struct {
	//
	// A Driver is an implementation of SchemaDescriptorBase which can digest the driver-specific manifest YAML
	// and also generate the object code for our solution.
	//
	Driver dataDriver.DataDriver `yaml:"driver"`
	//
	// The ToolChain is our output from the Driver's generator
	//
	ToolChain []dataCommon.ToolChainCodeBlocks
}