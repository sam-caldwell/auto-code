package manifest

type GlobalSection struct {
	Name string `yaml:"name"`

	Description string `yaml:"description"`

	Author string `yaml:"author"`

	Email string `yaml:"email"`

	Copyright string `yaml:"copyright"`

	License string `yaml:"license"`

	Language string `yaml:"language"`

	Version string `yaml:"version"`
}
