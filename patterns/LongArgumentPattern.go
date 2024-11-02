package patterns

const (

	// LongArgumentPattern - the "long" arguments (those starting with -- and followed by a longer identifier
	LongArgumentPattern = `^--[a-zA-Z][a-zA-Z0-9_\-\\.]{0,64}[a-zA-Z0-9]{1,}$`
)
