package arguments

import (
	"fmt"
	"github.com/sam-caldwell/ansi"
	"strings"
)

// PrintHelp - print the internal state
func (arg *Argument) PrintHelp() *Argument {
	if arg.err != nil {
		ansi.Red().Printf("error: %v\n", arg.err).Reset()
	}
	fmt.Println(arg.programName)
	fmt.Println(arg.copyright)
	fmt.Println(strings.Repeat("-", 78))
	for _, v := range arg.args {
		fmt.Printf(" %v\n", v.String())
	}
	return arg
}
