package manifest

// GlobalSection - Manifest Yaml global parameters section
type GlobalSection struct {
	//
	// Name - Represents program name
	Name NameIdentifier `yaml:"name"`
	//
	// Description - Represents a program description
	Description NonEmptyText `yaml:"description"`
	//
	// Author - Represents the program author
	Author NonEmptyText `yaml:"author"`
	//
	// Email - Represents the project email address
	Email EmailAddress `yaml:"email"`
	//
	// Copyright - Represents the project copyright
	Copyright CopyrightString `yaml:"copyright"`
	//
	// License - Represents the project license
	License LicenseString `yaml:"license"`
	//
	// Language - Represents the project language
	Language LanguageName `yaml:"language"`
	//
	// Version - Represents the project Version
	Version SemanticVersionString `yaml:"version"`
	//
	// GitRepoUrl - Represents the project code repo
	GitRepoUrl GitRepoUrlString `yaml:"git_repo,omitempty"`
}
