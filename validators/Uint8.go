package validator

// Uint8 - An unsigned integer Yaml Manifest Validator Parameter object
//
// This struct allows the consumer to parse the manifest yaml file and to
// verify the manifest file's minimum and maximum threshold values using
// the Verify() method.
type Uint8 struct {
	Min uint8 `yaml:"min"`
	Max uint8 `yaml:"max"`
}
