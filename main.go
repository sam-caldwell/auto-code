// main.go (Generator Entrypoint)
package main

import (
	"fmt"
	"github.com/sam-caldwell/ansi"
	"github.com/sam-caldwell/auto-code/generator"
	"github.com/sam-caldwell/auto-code/manifest"
	"github.com/sam-caldwell/auto-code/messages"
	"github.com/sam-caldwell/auto-code/misc"
)

const (
	errorExitCode = 1
)

func main() {

	var (
		projectManifest  manifest.Manifest
		manifestFileName string
		noop             bool
		codeRoot         string
	)
	// Get the command line arguments for the generator tool
	manifestFileName, codeRoot, noop = misc.GetCliArguments()

	if err := projectManifest.Load(manifestFileName); err != nil {
		ansi.Red().Printf(messages.ErrFailedToLoadManifest, err).Fatal(errorExitCode)
	}

	if err := generator.Run(&projectManifest, &codeRoot, noop); err != nil {
		ansi.Red().Printf(messages.ErrFailedToGenerateCode, err).Fatal(errorExitCode)
	}

	fmt.Println(messages.CodeGenerationCompleteOk)
}
