package manifest

// GlobalSection - Manifest Yaml global parameters section
type GlobalSection struct {
	Name NameIdentifier `yaml:"name"`

	Description NonEmptyText `yaml:"description"`

	Author NonEmptyText `yaml:"author"`

	Email EmailAddress `yaml:"email"`

	Copyright CopyrightString `yaml:"copyright"`

	License LicenseString `yaml:"license"`

	Language LanguageName `yaml:"language"`

	Version SemanticVersionString `yaml:"version"`

	GitRepoUrl GitRepoUrlString
}
