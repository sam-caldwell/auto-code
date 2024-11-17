package manifest

// String - returns the string representation of the UrlPattern
func (u *UrlPattern) String() string {

	result := u.scheme.String() + "://"

	if userInfo := u.userInfo.String(); userInfo != "" {
		result += userInfo + "@"
	}

	result += u.domain.String()

	if u.port != 0 {
		result += ":" + u.port.String()
	}

	result += u.path.String()
	if query := u.query; query != "" {
		result += "?" + u.query.String()
	}
	if fragment := u.fragment; fragment != "" {
		result += "#" + u.fragment.String()
	}
	return result
}
