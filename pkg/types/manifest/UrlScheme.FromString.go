package manifest

// FromString - store a given string as the UrlScheme
func (u *UrlScheme) FromString(s string) {
	*u = UrlScheme(s)
}
