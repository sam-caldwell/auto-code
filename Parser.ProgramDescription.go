package arguments

import (
	"fmt"
	"strings"
)

// ProgramDescription - Set the program's description (REQUIRED).
//
// A program should have a description which can be used in help text
// usage displays or other contexts to help users understand what the
// program does quickly and easily.
func (parser *Parser) ProgramDescription(description string) *Parser {

	if strings.TrimSpace(description) == "" {

		parser.err = fmt.Errorf("invalid program description")

	}

	parser.programDescription = description

	return parser

}
