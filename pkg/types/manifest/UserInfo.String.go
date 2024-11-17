package manifest

// String - return the RFC-3986 user info part as a string
func (u *UserInfo) String() string {
	result := ""
	if u.Username != "" {
		result = u.Username
		if u.Password != "" {
			result = result + ":" + u.Password
		}
	}
	return result
}
