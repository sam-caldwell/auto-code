package manifest

import (
	"fmt"
	"github.com/sam-caldwell/auto-code/manifest/messages"
	"github.com/sam-caldwell/auto-code/manifest/patterns"
	"regexp"
	"strings"
)

func (format *FileFormatString) Verify() error {
	*format = FileFormatString(strings.ToLower(strings.TrimSpace(string(*format))))
	if pattern := regexp.MustCompile(patterns.FileFormatPattern); pattern.MatchString(string(*format)) {
		return nil
	}
	return fmt.Errorf(messages.ErrInvalidFileFormat, *format)
}
