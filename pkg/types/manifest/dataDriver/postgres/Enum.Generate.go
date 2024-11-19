package postgres

import (
	"fmt"
)

// Generate - Generate the code for a postgresql enum
func (e *Enum) Generate() *CodeBlock {
	name := e.Name.String()
	elements := e.Elements.String()
	return &CodeBlock{
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

type CodeBlock struct {
	Forward  []string
	Rollback []string
}
