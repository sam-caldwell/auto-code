package arguments

import (
	"fmt"
	"strings"
)

// NewHelpText - initialize and return a new HelpText object
//
// This ensures there are no empty help text strings.  Every parameter should
// provide some idea of what the parameter is and how it is used.
func NewHelpText(helpText *string) (ht *HelpText, err error) {

	if text := strings.TrimSpace(*helpText); text == "" {

		*ht = ""
		err = fmt.Errorf(invalidHelpText)

	} else {
		*ht = HelpText(text)
		err = nil

	}

	return ht, err
}
