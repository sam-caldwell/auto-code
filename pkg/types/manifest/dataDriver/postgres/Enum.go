package postgres

import "strings"

// Enum - represents a database enumerated type
type Enum struct {
	Name     Identifier   `yaml:"name"`
	Elements EnumElements `yaml:"elements"`
}

type EnumElements []Identifier

func (e *EnumElements) String() string {
	s := make([]string, len(*e))
	for i, v := range *e {
		s[i] = string(v)
	}
	return strings.Join(s, ",")
}
