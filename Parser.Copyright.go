package arguments

import (
	"fmt"
	"regexp"
)

// Copyright - Set the copyright tag for an application (REQUIRED).
//
// Programs should have a defined and well-formatted copyright string
// to ensure license and intellectual property rights are properly
// asserted.
//
// This method tests for a valid US copyright string.  For example:
//
//	(c) 2024 Sam Caldwell.  All Rights Reserved.
//
// or
//
//	© 2024 Sam Caldwell.  All Rights Reserved.
func (parser *Parser) Copyright(copyright string) *Parser {

	pattern := regexp.MustCompile(`^(©|\(c\))?\s?(19|20)\d{2}(\s?,\s?(19|20)\d{2})*\s?[A-Za-z0-9\s,.'-]+$`)

	if pattern.MatchString(copyright) {

		parser.programCopyright = copyright

	} else {

		parser.programCopyright = ""
		parser.err = fmt.Errorf("invalid copyright string pattern")

	}

	return parser

}
