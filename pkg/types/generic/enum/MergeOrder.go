package enum

// MergeOrder - represents the order in which Configuration sections are merged
type MergeOrder uint8

const (
	Files MergeOrder = iota
	EnvironmentVars
	CommandLine
)
