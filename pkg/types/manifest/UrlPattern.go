package manifest

// UrlPattern - represents a valid RFC-3986 URL pattern
//
// Example:
//
//	protocol://user:password@domain:port/uri_path/?query=criteria
//
// where we may use ${identifier} as placeholders for configuration variables
// at any point in the pattern.
type UrlPattern struct {
	scheme   UrlScheme
	userInfo UserInfo
	domain   NetworkAddress
	port     NetworkPort
	path     UrlPath
	query    UrlQuery
	fragment UrlFragment
}
