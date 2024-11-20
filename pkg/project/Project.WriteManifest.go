package project

import (
	"github.com/sam-caldwell/ansi"
	"gopkg.in/yaml.v3"
)

// WriteManifest - write the resolved manifest to the target directory
func WriteManifest() error {
	result, err := yaml.Marshal(source)
	if err != nil {
		return err
	}
	ansi.Cyan().Println(string(result)).LF().Reset()
	//ToDo: write the resolved manifest to ${target_dir}/auto-code.yaml
	return nil
}
