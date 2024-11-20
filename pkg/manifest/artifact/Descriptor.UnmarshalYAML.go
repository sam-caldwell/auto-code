package artifact

import (
	"fmt"
	"github.com/sam-caldwell/auto-code/pkg/manifest"
	"gopkg.in/yaml.v3"
	"strings"
)

// UnmarshalYAML - Parse a Descriptor and store its information properly.
//
// This method will parse the YAML Descriptor and resolve any external $ref
// which may be included.
func (a *Descriptor) UnmarshalYAML(node *yaml.Node) error {

	// Parse the Descriptor, handling both $ref and non-$ref dataContract
	if err := manifest.ParseYamlObjectWithReferences(node, a); err != nil {
		return err
	}
	name := strings.ToLower(strings.TrimSpace(a.Name.String()))
	for _, dep := range a.Dependencies {
		if strings.ToLower(strings.TrimSpace(string(dep))) == name {
			return fmt.Errorf("circular reference: '%s' cannot depend upon itself", name)
		}
	}

}
