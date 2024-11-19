package manifest

import "github.com/sam-caldwell/auto-code/pkg/types/manifest/dataDriver"

// DataSchemaDescriptorBase - a base class used to create common functionality
//
// This is a 'base class' per se to define common functionality.
type DataSchemaDescriptorBase struct {
	Driver dataDriver.DataDriver `yaml:"driver"`
}
