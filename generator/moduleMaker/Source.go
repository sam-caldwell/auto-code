package moduleMaker

import (
	"github.com/sam-caldwell/types"
)

// Source - A descriptor identifying the template (input) and artifact (output)
type Source struct {
	template types.FileNameString
	artifact types.FileNameString
}
