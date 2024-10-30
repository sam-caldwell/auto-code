package arguments

import (
	"fmt"
	"strings"
)

// ProgramName - Set the program name for an application (REQUIRED).
//
// A program should have a name which can be used in help text and other
// contexts to help the user understand which program is running, etc.
// ToDo: Implement a process title rename such that the operating system
//
//	process table will show ProgramName as the name of a running
//	program's process.
func (parser *Parser) ProgramName(name string) *Parser {

	if strings.TrimSpace(name) == "" {
		parser.err = fmt.Errorf(invalidProgramName)
	}

	parser.programName = name

	return parser

}
