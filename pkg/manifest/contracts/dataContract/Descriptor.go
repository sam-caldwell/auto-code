package dataContract

import (
	"github.com/sam-caldwell/auto-code/pkg/manifest"
	"github.com/sam-caldwell/auto-code/pkg/manifest/artifact"
	"github.com/sam-caldwell/auto-code/pkg/manifest/contracts"
	"github.com/sam-caldwell/auto-code/pkg/types/generic"
)

// Descriptor - represents the content of a dataContract contract
type Descriptor struct {
	manifest.DataObjectWithReference
	Name      generic.Identifier          `yaml:"name"`
	Artifacts []artifact.Name             `yaml:"artifacts"`
	Schema    contracts.SchemaDescriptors `yaml:"schema"`
}
