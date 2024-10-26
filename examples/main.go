package main

import (
	"github.com/sam-caldwell/arguments"
)

func main() {
	args := arguments.New()
	args.
		ProgramName("Example Program").
		Copyright("(c) 2024 Sam Caldwell.  MIT License").
		//PositionalString(1,"pTest1","defaultValue1","a positional string1", `[a-z]{4}`).
		String("--test", "defaultString", "a test string", `[a-z]{3}`).
		//PositionalInt(3,"pTest3","defaultValue3","a positional string3", 1, 10).
		String("--foo", "fooValue", "foo must foo", arguments.NoValidation).
		Int("--number", 0, "test number", -1, 1).
		//PositionalString(2,"pTest2","defaultValue2","a positional string2", `[a-z]{4}`).
		Bool("--bool", true, "test boolean").
		Float("--float", 3.14, "test float", -1.0, 4.0).
		Flag("--flag", "this is a test flag.  It is true if present").
		PrintHelp()
}
