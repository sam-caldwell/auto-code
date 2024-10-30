Theory of Operation

```
package main

/*
 * We can define and parse arguments as a consistent
 * chain of function calls as follows:
 */
func getArgs() *arguments.Config {
    /*
     * First, we create a parser...
     */
	args := arguments.NewParser()
	/*
	 * Then we chain calls to Arguments (args) parser to
	 * both define and parse the command line arguments before
	 * reducing the result to a config object.
	 */
	config:=args.ProgramName("Example Program").
		Copyright("(c) 2024 Sam Caldwell.  MIT License").
		PositionalString(1,"pTest1","defaultValue1","a positional string1", `[a-z]{4}`).
		String("--test", "defaultString", "a test string", `[a-z]{3}`).
		PositionalInt(3,"pTest3","defaultValue3","a positional string3", 1, 10).
		String("--foo", "fooValue", "foo must foo", arguments.NoValidation).
		Int("--number", 0, "test number", -1, 1).
		PositionalString(2,"pTest2","defaultValue2","a positional string2", `[a-z]{4}`).
		Bool("--bool", true, "test boolean").
		Float("--float", 3.14, "test float", -1.0, 4.0).
		Flag("--flag", "this is a test flag.  It is true if present").
		PrintHelp().
		Parse()
    return config, args.Error()
}

func main(){
    /*
     * The resulting config object is a read-only data object
     * containing our configuration.
     */
    config:=getArgs()
}
```