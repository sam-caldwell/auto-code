package validator

// Float64 - A floating-point Yaml Manifest Validator Parameter object
//
// This struct allows the consumer to parse the manifest yaml file and to
// verify the manifest file's minimum and maximum threshold values using
// the Verify() method.
//
// A YAML Manifest Validator Parameter object is an implementation of the
// ParameterInterface type.
type Float64 struct {
	Min float64 `yaml:"min"`
	Max float64 `yaml:"max"`
}
