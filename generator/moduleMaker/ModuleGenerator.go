package moduleMaker

import "github.com/sam-caldwell/auto-code/manifest"

// ModuleGenerator - Generate common modules
type ModuleGenerator struct {
	projectManifest *manifest.Manifest
	moduleManifest  *ModuleManifest
	projectRoot     *string
}
