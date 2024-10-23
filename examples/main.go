package main

import (
	"github.com/sam-caldwell/arguments"
)

func main() {
	args := arguments.New()
	args.
		ProgramName("Example Program").
		Copyright("(c) 2024 Sam Caldwell.  MIT License").
		String("--test", "defaultString", "a test string", `[a-z]{3}`).
		String("--foo", "fooValue", "foo must foo", arguments.NoValidation).
		Int("--number", 0, "test number", -1, 1).
		Bool("--bool", true, "test boolean").
		Float("--float", 3.14, "test float", -1.0, 4.0).
		Flag("--flag", "this is a test flag.  It is true if present").
		PrintHelp()
}
