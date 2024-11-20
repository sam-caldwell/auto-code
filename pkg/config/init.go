package config

import (
	"flag"
	"github.com/sam-caldwell/ansi"
	"github.com/sam-caldwell/directory"
	"github.com/sam-caldwell/file"
)

// init - initialize the configuration state at startup
func init() {

	manifestFile := flag.String("manifest-file", "manifest.yaml", "Path to the manifest file")
	targetDirectory := flag.String("target-directory", "build", "Path to the target directory")
	debug := flag.Bool("debug", false, "Enable debug mode")
	flag.Parse()

	if !file.Exists(*manifestFile) {
		ansi.Red().Printf("manifest file not found (%s)", *manifestFile).LF().Fatal(1)
	}

	if !directory.Exists(*targetDirectory) {
		ansi.Red().Printf("targetDirectory not found (%s)", *targetDirectory).LF().Fatal(1)
	}

	config.ManifestFile = *manifestFile
	config.TargetDirectory = *targetDirectory
	config.Debug = *debug
}
