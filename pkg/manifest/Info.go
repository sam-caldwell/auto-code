package manifest

type Info struct {
	Title       NonEmptyString  `yaml:"title"`
	Description NonEmptyString  `yaml:"description"`
	Author      NonEmptyString  `yaml:"author"`
	AuthorEmail EmailAddress    `yaml:"authorEmail"`
	Copyright   USCopyright     `yaml:"copyright"`
	License     LicenseName     `yaml:"license"`
	Version     SemanticVersion `yaml:"version"`
}
