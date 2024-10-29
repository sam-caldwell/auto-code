package arguments

import "fmt"

// storeArgument - Extract, validate and store an expected (detected) argument.
//
// When this is called, we have a token which starts with '-' or '--' and we expect an argument.
func (arg *Argument) storeArgument(token *string, expectValue *bool, expectArgument *string, currentClass *ArgumentClass) {
	actualArgument, tokenIsValid := arg.args[*token]
	if !tokenIsValid {
		arg.err = fmt.Errorf(unknownOrUnexpectedArgument)
		return
	}
	// We have confirmed that this arg (starting with - or --) exists in our arguments map.
	switch class := actualArgument.class; class {
	case Bool, Directory, Email, File, Float, Int, Secret, String, Uint:
		*currentClass = class
		*expectArgument = *token
		*expectValue = true
		// happy path: we will need to get a value in the next cycle.
	case Flag:
		//thisArg is a flag.  There is no value.
		*currentClass = Flag
		actualArgument.value = true          // a present flag is a true flag
		*expectArgument = NoArgumentExpected // we do not expect an argument
		*expectValue = false
	default:
		arg.err = fmt.Errorf(unsupportedArgumentClass)
		break
	}
}
