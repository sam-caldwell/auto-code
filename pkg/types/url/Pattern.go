package url

import (
	"github.com/sam-caldwell/auto-code/pkg/manifest"
	"github.com/sam-caldwell/auto-code/pkg/types/network"
)

// Pattern - represents a valid RFC-3986 URL pattern
//
// Example:
//
//	protocol://user:password@domain:port/uri_path/?query=criteria
//
// where we may use ${identifier} as placeholders for configuration variables
// at any point in the pattern.
type Pattern struct {
	scheme   Scheme
	userInfo manifest.UserInfo
	domain   network.Address
	port     network.Port
	path     Path
	query    Query
	fragment Fragment
}
