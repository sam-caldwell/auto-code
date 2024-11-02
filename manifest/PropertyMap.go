package manifest

import (
	data2 "github.com/sam-caldwell/auto-code/data"
)

// PropertyMap - a map associating a dot-delimited object property path with a config property.
type PropertyMap map[data2.NameIdentifier]data2.ObjectPropertyString
