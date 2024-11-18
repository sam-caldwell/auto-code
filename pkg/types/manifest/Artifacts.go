package manifest

// Artifacts - represents the collection of artifacts in a solution
//
// Note: the artifacts object does NOT support $ref, though its children do.
type Artifacts struct {
	Artifact []ArtifactDescriptor
}
