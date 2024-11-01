// main.go (Generator Entrypoint)
package main

import (
	"flag"
	"fmt"
	"github.com/sam-caldwell/ansi"
	"github.com/sam-caldwell/auto-code/generator"
	"github.com/sam-caldwell/auto-code/manifest"
	"github.com/sam-caldwell/file"
)

const (
	defaultManifest   = "manifest.yaml"
	manifestFlagUsage = "Identify the project descriptor file"
	manifestFlag      = "--manifest"
	defaultNoop       = false
	noopUsage         = "do not generate code artifacts"
	noopFlag          = "--noop"
	codeRootFlag      = "--projectRoot"
	defaultCodeRoot   = "./"
	codeRootUsage     = "the root directory for the project (where the generated code will be written)"
	errorExitCode     = 1
)

func main() {
	var projectManifest manifest.Manifest
	manifestFileName := flag.String(manifestFlag, defaultManifest, manifestFlagUsage)
	noop := flag.Bool(noopFlag, defaultNoop, noopUsage)
	codeRoot := flag.String(codeRootFlag, defaultCodeRoot, codeRootUsage)
	flag.Parse()

	if !file.Exists(*manifestFileName) {
		ansi.Red().Printf("manifest file (%s) not found\n", *manifestFileName).Fatal(errorExitCode)
	}

	if err := projectManifest.Load("manifest.yaml"); err != nil {
		ansi.Red().Printf("Failed to load manifest: '%s'\n", err).Fatal(errorExitCode)
	}

	if err := generator.Run(&projectManifest, codeRoot, *noop); err != nil {
		ansi.Red().Printf("Failed to generate code: '%s'\n", err).Fatal(errorExitCode)
	}

	fmt.Println("Code generation complete!")
}
