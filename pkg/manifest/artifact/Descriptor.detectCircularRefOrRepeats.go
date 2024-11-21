package artifact

import (
	"fmt"
	"github.com/sam-caldwell/auto-code/pkg/types/generic"
	"strings"
)

// detectCircularRefOrRepeats - Make sure we do not have circular dependencies or repeating dependencies
func (a *Descriptor) detectCircularRefOrRepeats() error {

	seenBefore := make(map[generic.Identifier]struct{})

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
