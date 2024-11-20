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

	if err := a.Name.Validate(); err != nil {
		return err
	}

	return a.detectCircularRefOrRepeats()

}

func (a *Descriptor) detectCircularRefOrRepeats() error {
	// Make sure we do not have circular dependencies or repeating dependencies
	seenBefore := make(map[Name]struct{})
	for _, dep := range a.Dependencies {
		if strings.ToLower(strings.TrimSpace(string(dep))) == a.Name.String() {
			return fmt.Errorf("circular reference: '%s' cannot depend upon itself", a.Name.String())
		}
		if _, ok := seenBefore[dep]; ok {
			return fmt.Errorf("dependencies can only appear once: '%s'", dep.String())
		}
		seenBefore[dep] = struct{}{}
	}
	return nil
}
