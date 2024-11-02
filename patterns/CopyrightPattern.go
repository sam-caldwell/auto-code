package patterns

const (
	// CopyrightPattern - copyright string pattern
	//
	// All copyrights must follow the US pattern:
	//   (c) YEAR Author. All Rights Reserved.
	// or
	//   ©YEAR Author. All Rights Reserved.
	CopyrightPattern = `^©\s?(\d{4}(-\d{4})?)(,\s?\d{4}(-\d{4})?)*\s[A-Za-z .,'-]+$`
)
