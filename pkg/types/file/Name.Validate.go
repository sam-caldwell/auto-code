package file

import (
	"fmt"
	"regexp"
	"strings"
)

// Validate - validate the file name
func (f *Name) Validate() error {

	const pattern = `^(/?([a-zA-Z0-9_.+-]+|[!#$&'()*+,;=@^{}|~]+)(/[a-zA-Z0-9_.+-]+|[!#$&'()*+,;=@^{}|~]+)*)?$`

	trimmedValue := strings.TrimSpace(f.String())

	if re := regexp.MustCompile(pattern); !re.MatchString(trimmedValue) {
		return fmt.Errorf("invalid path/file name (%s)", trimmedValue)
	}

	*f = Name(trimmedValue)

	return nil

}
