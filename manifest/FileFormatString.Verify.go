package manifest

import (
	"fmt"
	"regexp"
	"strings"
)

func (format *FileFormatString) Verify() error {
	*format = FileFormatString(strings.ToLower(strings.TrimSpace(string(*format))))
	if pattern := regexp.MustCompile(fileFormatPattern); pattern.MatchString(string(*format)) {
		return nil
	}
	return fmt.Errorf(errInvalidFileFormat, *format)
}
