package dataCommon

// ToolChainCodeBlocks - The 'base class' for defining code blocks in our data-layer toolchain.
//
// Each specific data driver should define this map.
type ToolChainCodeBlocks struct {
	Features map[Identifier]CodeBlock //tools in the tool kit.
}
