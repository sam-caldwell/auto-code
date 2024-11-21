package artifact

import (
	"github.com/sam-caldwell/ansi"
	"github.com/sam-caldwell/auto-code/pkg/types/xref"
	"gopkg.in/yaml.v3"
)

// UnmarshalYAML - Parse a Descriptor and store its information properly.
//
// This method will parse the YAML Descriptor and resolve any external $ref
// which may be included.
func (a *Descriptor) UnmarshalYAML(node *yaml.Node) error {

	ansi.Debugln("artifact/Descriptor.UnmarshalYAML start").LF()
	defer ansi.Debugln("artifact/Descriptor.UnmarshalYAML done").LF()

	// Parse the Descriptor, handling both $ref and non-$ref dataContract
	if err := xref.ParseYamlObjectWithReferences(node, a); err != nil {
		return err
	}

	if err := a.Name.Validate(); err != nil {
		return err
	}

	return a.detectCircularRefOrRepeats()

}
