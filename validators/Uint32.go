package validator

// Uint32 - An unsigned integer Yaml Manifest Validator Parameter object
//
// This struct allows the consumer to parse the manifest yaml file and to
// verify the manifest file's minimum and maximum threshold values using
// the Verify() method.
//
// A YAML Manifest Validator Parameter object is an implementation of the Interface type.
type Uint32 struct {
	Min uint32 `yaml:"min"`
	Max uint32 `yaml:"max"`
}
