package manifest

// String - return a string representation of an RFC-3986 scheme
func (u *UrlScheme) String() string {
	return string(*u)
}
