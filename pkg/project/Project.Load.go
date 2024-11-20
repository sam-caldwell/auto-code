package project

import (
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
	//Load/Parse/Resolve the project manifest
	var raw []byte
	raw, err = os.ReadFile(config.Manifest())
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(raw, &source); err != nil {
		return err
	}
	return nil
}
