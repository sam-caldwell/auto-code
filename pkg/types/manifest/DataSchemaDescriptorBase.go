package manifest

// DataSchemaDescriptorBase - a base class used to create common functionality
//
// This is a 'base class' per se to define common functionality.
type DataSchemaDescriptorBase struct {
	Driver DataDriver `yaml:"driver"`
}
