package common

// CodeBlock - represents the forward and rollback code sequences for a given feature
type CodeBlock struct {
	Forward  []string
	Rollback []string
}
