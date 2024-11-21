package metadata

import (
	"github.com/sam-caldwell/ansi"
	"gopkg.in/yaml.v3"
)

// UnmarshalYAML - unmarshal Info metadata
func (sv *Info) UnmarshalYAML(node *yaml.Node) error {

	ansi.Debugln("unmarshal Info: start").LF()

	if err := node.Decode(sv); err != nil {
		return err
	}

	ansi.Debugln("unmarshal Info: done").LF()
	return nil
}
