// Package generator generator/Run.go
package generator

import (
	generatorTemplate "github.com/sam-caldwell/auto-code/generator/templates"
	"github.com/sam-caldwell/auto-code/manifest"
	"os"
	"text/template"
)

// Run - generate code for a project given the manifest.
//
// This is the entrypoint for the code generator.  It should be kept small and clean
// and call other mission-specific parts.
func Run(manifest *manifest.Manifest, projectRoot *string, noop bool) (err error) {
	if err := manifest.SetGitRepo(); err != nil {
		return err
	}
	const (
		gitRepoRoot = "github.com/sam-caldwell/"
	)
	// Create a list of module generators (packages in some languages).  This should be defined in the
	// order in which they will be generated.
	moduleList := []ModuleGeneratorList{
		newModuleGenerator(manifest, generatorTemplate.Configuration, "configuration", "Configuration"),
		// ToDo: add more modules
	}
	for _, module := range moduleList {
		module.Generate(projectRoot, noop)
	}

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
