package common

import "github.com/sam-caldwell/auto-code/pkg/types/generic"

// ToolChainCodeBlocks - The 'base class' for defining code blocks in our dataContract-layer toolchain.
//
// Each specific dataContract driver should define this map.
type ToolChainCodeBlocks struct {
	Features map[generic.Identifier]CodeBlock //tools in the tool kit.
}
