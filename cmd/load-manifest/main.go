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
		ansi.Red().Errorf("Error loading manifest: %s", err).LF().Fatal(1)
	}
	ansi.Debugf("Writing manifest...").LF()
	if err := project.WriteManifest(); err != nil {
		ansi.Red().Errorf("Error writing manifest: %s", err).LF().Fatal(2)
	}
	ansi.Debugf("Done.").LF().Reset()
}
