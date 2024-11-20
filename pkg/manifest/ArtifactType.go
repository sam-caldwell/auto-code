package manifest

type ArtifactType int

const (
	service ArtifactType = iota
	external
	binary
)
