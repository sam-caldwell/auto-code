package validator

// Float32 - A floating-point Yaml Manifest Validator Parameter object
//
// This struct allows the consumer to parse the manifest yaml file and to
// verify the manifest file's minimum and maximum threshold values using
// the Verify() method.
//
// A YAML Manifest Validator Parameter object is an implementation of the
// Interface type.
type Float32 struct {
	Min float32 `yaml:"min"`
	Max float32 `yaml:"max"`
}
