package manifest

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"regexp"
	"strings"
)

// UnmarshalYAML - unmarshal a YAML object into the FileName string
//
// This should accept a
func (f *FileName) UnmarshalYAML(node *yaml.Node) error {

	const pattern = `^(/?([a-zA-Z0-9_.+-]+|[!#$&'()*+,;=@^{}|~]+)(/[a-zA-Z0-9_.+-]+|[!#$&'()*+,;=@^{}|~]+)*)?$`

	var value string

	if err := node.Decode(&value); err != nil {
		return err
	}

	trimmedValue := strings.TrimSpace(value)

	if re := regexp.MustCompile(pattern); !re.MatchString(trimmedValue) {
		return fmt.Errorf("invalid path/file name (%s)", trimmedValue)
	}

	*f = FileName(value)

	return nil

}
