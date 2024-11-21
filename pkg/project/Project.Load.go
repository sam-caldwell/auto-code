package project

import (
	"github.com/sam-caldwell/ansi"
	"github.com/sam-caldwell/auto-code/pkg/config"
	"gopkg.in/yaml.v3"
	"os"
)

// Load - Load/Parse/Resolve the project manifest.
//
// This will load the YAML file(s) and resolve any $ref references.
// The YAML will be parsed and resolved to a final state.  When this
// is completed, the project.source variable will contain the entire
// project structure in memory for final generation.
func Load() (err error) {
	var raw []byte
	ansi.Debugf("Reading file (%s)", config.Manifest()).LF()
	raw, err = os.ReadFile(config.Manifest())
	if err != nil {
		return err
	}
	ansi.Debugln("Unmarshal YAML manifest")
	if err = yaml.Unmarshal(raw, &source); err != nil {
		return err
	}
	ansi.Debugln("manifest is loaded successfully")
	return nil
}
