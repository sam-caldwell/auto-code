package postgres

import (
	"github.com/sam-caldwell/auto-code/pkg/types/generic"
)

// Enum - represents a database enumerated type
type Enum struct {
	Name     generic.Identifier `yaml:"name"`
	Elements EnumElements       `yaml:"elements"`
}

type EnumElements []generic.Identifier

func (e *EnumElements) StringArray() []string {
	s := make([]string, len(*e))
	for i, v := range *e {
		s[i] = string(v)
	}
	return s
}
