package manifest

func (t *ArtifactType) String() string {
	return [...]string{"service", "external", "binary"}[*t]
}
