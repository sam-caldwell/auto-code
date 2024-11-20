package project

import (
	"github.com/sam-caldwell/ansi"
	"github.com/sam-caldwell/auto-code/pkg/config"
)

// Generate - generate the final project artifacts.
//
// Given the loaded 'project.source' produced by the Load()
// function, the Generate() function will iterate over the
// artifacts and generate the source code for each in order.
func Generate() (err error) {

	for order, artifact := range source.Artifacts {
		ansi.Blue().
			Printf("artifact #%d (%s) generating...", order, artifact.Name.String()).
			LF().
			Reset()
		if err := artifact.Generate(config.Target(), config.Debug()); err != nil {
			return err
		}
	}

	return nil
}
