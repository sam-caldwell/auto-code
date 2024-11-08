package misc

import (
	"flag"
	"github.com/sam-caldwell/ansi"
	"github.com/sam-caldwell/auto-code/messages"
	"github.com/sam-caldwell/directory"
	"github.com/sam-caldwell/file"
)

const (
	defaultManifest   = "manifest.yaml"
	manifestFlagUsage = "Identify the project descriptor file"
	manifestFlag      = "--manifest"

	defaultNoop = false
	noopUsage   = "do not generate code artifacts"
	noopFlag    = "--noop"

	codeRootFlag    = "--projectRoot"
	defaultCodeRoot = "./"
	codeRootUsage   = "the root directory for the project (where the generated code will be written)"

	errorExitCode = 1
)

// GetCliArguments - Get commandline arguments
func GetCliArguments() (manifestFileName, codeRoot string, noop bool) {
	manifestFileNameFlag := flag.String(manifestFlag, defaultManifest, manifestFlagUsage)
	noopFlag := flag.Bool(noopFlag, defaultNoop, noopUsage)
	codeRootFlag := flag.String(codeRootFlag, defaultCodeRoot, codeRootUsage)
	flag.Parse()

	if !file.Exists(*manifestFileNameFlag) {
		ansi.Red().Printf(messages.ErrManifestFileNotFound, *manifestFileNameFlag).Fatal(errorExitCode)
	}

	if !directory.Exists(*codeRootFlag) {
		ansi.Red().Printf(messages.ErrTargetDirectoryExists, *codeRootFlag).Fatal(errorExitCode)
	}
	return *manifestFileNameFlag, *codeRootFlag, *noopFlag
}
