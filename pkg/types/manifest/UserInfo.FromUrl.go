package manifest

import "net/url"

// FromUrl - parse/validate username/password from RFC3986 user information
func (u *UserInfo) FromUrl(parsedUrl *url.URL) {

	if userInfo := parsedUrl.User; userInfo != nil {
		u.Username = userInfo.Username()
		if pass, isSet := userInfo.Password(); isSet {
			u.Password = pass
		}
	}
}
