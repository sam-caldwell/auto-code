// Package generator generator/generator.go
package generator

import (
	"github.com/sam-caldwell/auto-code/manifest"
	"os"
	"text/template"
)

// GenerateCode generates Go code based on the manifest settings
func GenerateCode(manifest *manifest.Manifest) error {

	// Set up the template for main.go
	mainTemplate, err := template.ParseFiles("templates/main.go")
	if err != nil {
		return err
	}

	// Create the main.go file
	file, err := os.Create("output/main.go")
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	// Populate the main template with manifest data
	if err := mainTemplate.Execute(file, manifest); err != nil {
		return err
	}

	return nil
}
