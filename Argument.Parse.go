package arguments

import (
	"fmt"
	"os"
)

func (arg *Argument) Parse() *Argument {

	var expectValue bool = false
	var expectArgument string
	var currentClass ArgumentClass
	var position uint //A position which ignores any optional arguments

	for _, token := range os.Args[1:] { // iterate over all arguments other than Args[0] (program name)
		if expectValue {
			arg.storeValue(&token, &expectValue, &expectArgument, &currentClass)
			continue
		}
		// expect -arg or --argument or positional argument
		if detectOptionalArgument(&token) {
			if expectValue {
				// error state: we expect a value or no argument but detected an argument
				arg.err = fmt.Errorf(unknownOrUnexpectedArgument)
			} else {
				arg.storeArgument(&token, &expectValue, &expectArgument, &currentClass)
			}
		} else {
			// we encountered a non-argument token.  This could be a positional argument.
			arg.storePositionalArgument(&token, position, &expectValue, &expectArgument, &currentClass)
			position++
		}
		if arg.err != nil {
			break
		}
	}
	return arg
}
