package validator

// Uint16 - An unsigned integer Yaml Manifest Validator Parameter object
//
// This struct allows the consumer to parse the manifest yaml file and to
// verify the manifest file's minimum and maximum threshold values using
// the Verify() method.
type Uint16 struct {
	Min uint16 `yaml:"min"`
	Max uint16 `yaml:"max"`
}
