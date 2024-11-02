package generator

import (
	"github.com/sam-caldwell/auto-code/data"
	generatorTemplate "github.com/sam-caldwell/auto-code/generator/templates"
	"github.com/sam-caldwell/auto-code/manifest"
	"github.com/sam-caldwell/types"
)

func newModuleGenerator(projectManifest *manifest.Manifest) *ModuleGenerator {
	return &ModuleGenerator{
		projectManifest: projectManifest,
		moduleManifest:  NewModuleManifest(&projectManifest.Global.Language),
	}
}

func NewModuleManifest(language *data.LanguageName) *ModuleManifestDescriptor {
	return &(generatorTemplate.ModuleManifest[language])
}

// ModuleGenerator - Generate common modules
type ModuleGenerator struct {
	projectManifest *manifest.Manifest
	moduleManifest  *ModuleManifest
	projectRoot     *string
}

type Source struct {
	template types.FileNameString
	artifact types.FileNameString
}

type ModuleName string

type ModuleDescriptor struct {
	Enabled bool
	files   []Source
}

type ModuleManifest map[ModuleName]ModuleDescriptor
