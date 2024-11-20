package url

import (
	"fmt"
	"github.com/sam-caldwell/auto-code/pkg/types/network"
	"gopkg.in/yaml.v3"
	"net/url"
	"strings"
)

// UnmarshalYAML - unmarshal and parse a parameterized RFC-3986 URL pattern from YAML
func (u *Pattern) UnmarshalYAML(node *yaml.Node) error {
	var rawUrl string
	if err := node.Decode(&rawUrl); err != nil {
		return err
	}

	// Parse the URL using the net/url package
	parsedUrl, err := url.Parse(strings.TrimSpace(rawUrl))
	if err != nil {
		return fmt.Errorf("failed to parse URL: %w", err)
	}

	u.scheme.FromString(parsedUrl.Scheme)
	u.userInfo.FromUrl(parsedUrl)

	if host := parsedUrl.Host; strings.Contains(host, ":") {
		hostParts := strings.Split(host, ":")
		u.domain.FromString(hostParts[0])
		if err := u.port.FromString(hostParts[1]); err != nil {
			return err
		}
	} else {
		u.domain = network.Address(host)
	}

	u.path.FromString(parsedUrl.Path)
	u.query.FromString(parsedUrl.RawQuery)
	u.fragment.FromString(parsedUrl.Fragment)

	return nil
}
