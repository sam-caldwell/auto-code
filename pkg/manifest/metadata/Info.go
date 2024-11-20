package manifest

import (
	"github.com/sam-caldwell/auto-code/pkg/types/copyright"
	"github.com/sam-caldwell/auto-code/pkg/types/email"
	"github.com/sam-caldwell/auto-code/pkg/types/generic"
	"github.com/sam-caldwell/auto-code/pkg/types/license"
	"github.com/sam-caldwell/auto-code/pkg/types/version"
)

type Info struct {
	Title       generic.NonEmptyString `yaml:"title"`
	Description generic.NonEmptyString `yaml:"description"`
	Author      generic.NonEmptyString `yaml:"author"`
	AuthorEmail email.Address          `yaml:"authorEmail"`
	Copyright   copyright.US           `yaml:"copyright"`
	License     license.Name           `yaml:"license"`
	Version     version.Version        `yaml:"version"`
}
