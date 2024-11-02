package validator

// Pattern - A string pattern (regex) YAML Manifest Validator Parameter object
//
// This struct allows the consumer to parse the manifest yaml file and to
// verify the manifest file's regular expression threshold values using
// the Verify() method.
//
// A YAML Manifest Validator Parameter object is an implementation of the Interface type.
type Pattern struct {
	Regex string `yaml:"regex"`
	Match bool   `yaml:"match"`
}
