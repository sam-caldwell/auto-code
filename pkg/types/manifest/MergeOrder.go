package manifest

// MergeOrder - represents the order in which Configuration sections are merged
type MergeOrder int

const (
	Files MergeOrder = iota
	EnvironmentVars
	CommandLine
)