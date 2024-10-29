package arguments

import (
	"strings"
)

func detectArgument(token *string) bool {
	const (
		shortArg = "-"
		longArg  = "--"
	)
	s := strings.TrimSpace(*token)
	return (strings.HasPrefix(s, longArg) || strings.HasPrefix(s, shortArg)) && s != longArg && s != shortArg
}
