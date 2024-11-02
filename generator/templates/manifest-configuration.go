package templates

import "github.com/sam-caldwell/auto-code/generator"

// Configuration - a map of language-specific yaml manifests used to define the Configuration module
var Configuration = generator.ModuleManifest{
	"c++":generator.ModuleDescriptor{
		Enabled: false,
		files:[]Source{},
	},
	"go":generator.ModuleDescriptor{
		Enabled: true,
		files:[]Source{},
	},
	"node":generator.ModuleDescriptor{
		Enabled: false,
		files:[]Source{},
	},
}
	"c++": `---
enabled: false
files:
  - template: Configuration.tmpl.h
    artifact: configuration/Configuration.h
`,
	"go": `---
enabled: true
files:
  - template: Configuration.Get.tmpl.go
    artifact: configuration/Configuration.Get.go
  - template: Configuration.Get.tmpl.go
    artifact: configuration/Configuration.Get.go
`,
	"node": `---
enabled: false
files:
  - template: Configuration.tmpl.js
    artifact: configuration/Configuration.js
`,
}
