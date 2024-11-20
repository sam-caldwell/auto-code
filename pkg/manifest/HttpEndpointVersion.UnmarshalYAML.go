package manifest

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// UnmarshalYAML - Unmarshal a yaml object to HttpEndpointVersion
func (h *HttpEndpointVersion) UnmarshalYAML(node *yaml.Node) error {

	const pattern = `^v[0-9]{1,7}$`

	var value string

	if err := node.Decode(&value); err != nil {
		return err
	}

	trimmedValue := strings.ToLower(strings.TrimSpace(value))

	if re := regexp.MustCompile(pattern); !re.MatchString(trimmedValue) {
		return fmt.Errorf("invalid API version (%s)", trimmedValue)
	}

	trimmedValue = strings.TrimPrefix(trimmedValue, "v")

	versionNumber, err := strconv.ParseInt(trimmedValue, 10, 64)

	if err != nil {
		return err
	}

	if mx := int64(math.MaxUint); versionNumber > mx {
		return fmt.Errorf("version number too large (%d) and must be less than (%d)", versionNumber, mx)
	}

	*h = HttpEndpointVersion(versionNumber)

	return nil

}
