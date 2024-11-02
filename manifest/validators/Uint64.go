package validator

// Uint64 - An unsigned integer Yaml Manifest Validator Parameter object
//
// This struct allows the consumer to parse the manifest yaml file and to
// verify the manifest file's minimum and maximum threshold values using
// the Verify() method.
type Uint64 struct {
	Min uint64 `yaml:"min"`
	Max uint64 `yaml:"max"`
}
