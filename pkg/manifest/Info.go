package manifest

import (
	"github.com/sam-caldwell/auto-code/pkg/types/license"
)

type Info struct {
	Title       NonEmptyString  `yaml:"title"`
	Description NonEmptyString  `yaml:"description"`
	Author      NonEmptyString  `yaml:"author"`
	AuthorEmail EmailAddress    `yaml:"authorEmail"`
	Copyright   USCopyright     `yaml:"copyright"`
	License     license.Name    `yaml:"license"`
	Version     SemanticVersion `yaml:"version"`
}
