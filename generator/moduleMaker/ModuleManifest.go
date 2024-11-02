package moduleMaker

import (
	"github.com/sam-caldwell/auto-code/data"
	"github.com/sam-caldwell/auto-code/manifest"
)

// ModuleManifest - a set of ModuleDescriptors, identified by language
type ModuleManifest map[data.LanguageName]ModuleDescriptor

func newModuleGenerator(projectManifest *manifest.Manifest) *ModuleGenerator {
	return &ModuleGenerator{
		projectManifest: projectManifest,
		moduleManifest:  NewModuleManifest(&projectManifest.Global.Language),
	}
}

func NewModuleManifest(language *data.LanguageName) *ModuleManifestDescriptor {
	return &(generatorTemplate.ModuleManifest[language])
}
