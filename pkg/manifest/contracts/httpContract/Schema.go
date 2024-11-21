package httpContract

import (
	"github.com/sam-caldwell/auto-code/pkg/manifest/contracts/httpContract/httpDriver"
	"github.com/sam-caldwell/auto-code/pkg/types/xref"
)

// Schema - defines the http endpoint schema
type Schema struct {
	xref.DataObjectWithReference
	//
	// A Driver is an implementation of SchemaDescriptorBase which can digest the driver-specific manifest YAML
	// and also generate the object code for our solution.
	//
	Driver httpDriver.HttpDriver `yaml:"driver"`
}
