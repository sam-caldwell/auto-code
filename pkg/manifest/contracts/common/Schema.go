package common

import "gopkg.in/yaml.v3"

type Schema interface {
	Unmarshal(node *yaml.Node) error
	Generate() (error, CodeBlock)
}
