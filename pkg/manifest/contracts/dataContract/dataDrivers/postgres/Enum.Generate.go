package postgres

import (
	"fmt"
	"github.com/sam-caldwell/auto-code/pkg/manifest/contracts/common"
)

// Generate - Generate the code for a postgresql enum
func (e *Enum) Generate() *common.CodeBlock {

	name := e.Name.String()
	elements := e.Elements.StringArray()

	return &common.CodeBlock{

		Forward: []string{
			// Note: upsertEnum() is a function auto-code must create before executing migrations.
			//       This is part of our toolkit.  upsertType should allow create or update of an Enum.
			fmt.Sprintf("SELECT upsertEnum('%s','%s');", name, elements),
		},

		Rollback: []string{
			fmt.Sprintf("DROP TYPE IF EXISTS '%s'", name),
		},
	}
}
