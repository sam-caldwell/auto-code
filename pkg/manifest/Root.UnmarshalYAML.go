package manifest

import (
	"fmt"
	"github.com/sam-caldwell/ansi"
	"gopkg.in/yaml.v3"
)

// UnmarshalYAML - unmarshal the manifest from the root
func (root *Root) UnmarshalYAML(node *yaml.Node) error {

	i := len(node.Content) - 1
	ansi.Debugf("unmarshal yaml at root starting. node {line: %d, value: '%v'}",
		node.Line, node.Content[i]).LF()

	if err := node.Decode(&root); err != nil {
		return err
	}

	if root.Version.String() != "v0.0.1" {
		return fmt.Errorf("unsupported auto-code version %s (expect v0.0.1)", root.Version.String())
	}

	ansi.Debugln("unmarshal yaml at root success.")

	return nil
}
