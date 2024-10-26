package arguments

import "strings"

func detectOptionalArgument(token *string) bool {
	return strings.HasPrefix(*token, "-") || strings.HasPrefix(*token, "--")
}
