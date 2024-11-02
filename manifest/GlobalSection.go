package manifest

import (
	data2 "github.com/sam-caldwell/auto-code/data"
	"github.com/sam-caldwell/auto-code/gitrepo"
	"github.com/sam-caldwell/auto-code/manifest/copyright"
)

// GlobalSection - Manifest Yaml global parameters section
type GlobalSection struct {
	//
	// Name - Represents program name
	Name data2.NameIdentifier `yaml:"name"`
	//
	// Description - Represents a program description
	Description data2.NonEmptyText `yaml:"description"`
	//
	// Author - Represents the program author
	Author data2.NonEmptyText `yaml:"author"`
	//
	// Email - Represents the project email address
	Email data2.EmailAddress `yaml:"email"`
	//
	// Copyright - Represents the project copyright
	Copyright copyright.String `yaml:"copyright"`
	//
	// License - Represents the project license
	License data2.LicenseString `yaml:"license"`
	//
	// Language - Represents the project language
	Language data2.LanguageName `yaml:"language"`
	//
	// Version - Represents the project Version
	Version data2.SemanticVersionString `yaml:"version"`
	//
	// GitRepoUrl - Represents the project code repo
	GitRepoUrl gitrepo.UrlString `yaml:"git_repo,omitempty"`
}
