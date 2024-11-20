package manifest

type HttpContract struct {
	Artifacts []ArtifactName           `yaml:"artifacts"`
	Endpoints []HttpEndpointDescriptor `yaml:"endpoints"`
}
