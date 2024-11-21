package main

import (
	"github.com/sam-caldwell/ansi"
	"github.com/sam-caldwell/auto-code/pkg/config"
	"github.com/sam-caldwell/auto-code/pkg/project"
)

func main() {
	if config.Debug() {
		ansi.EnableDebug()
	}
	ansi.Debugf("start program...").LF()
	if err := project.Load(); err != nil {
		ansi.Red().Printf("Error loading manifest: %s", err).LF().Reset()
	}
	ansi.Debugf("Writing manifest...").LF()
	if err := project.WriteManifest(); err != nil {
		ansi.Red().Printf("Error writing manifest: %s", err).LF().Reset()
	}
	ansi.Debugf("Generating solution...").LF()
	if err := project.Generate(); err != nil {
		ansi.Red().Printf("Error generating project: %s", err).LF().Reset()
	}
	ansi.Debugf("Done.").LF().Reset()
}
