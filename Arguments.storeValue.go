package arguments

import "fmt"

// storeValue - Store a token as a value if it is not an argument since we expect a value.
func (arg *Arguments) storeValue(token *string, expectValue *bool, expectArgument *string, currentClass *ArgumentClass) {
	if detectArgument(token) {
		arg.err = fmt.Errorf(unknownOrUnexpectedArgument)
		return
	}
	switch currentArgument := arg.args[*expectArgument]; currentArgument.class {
	case Bool:
		arg.err = arg.args[*expectArgument].SetValueBool(token)
	case Flag:
		trueString := "true"
		arg.err = arg.args[*expectArgument].SetValueBool(&trueString)
	case Float:
		arg.err = arg.args[*expectArgument].SetValueFloat(token)
	case Int:
		arg.err = arg.args[*expectArgument].SetValueInt(token)
	case String, Directory, File, Secret, Email:
		arg.err = arg.args[*expectArgument].SetValueString(token)
	case Uint:
		arg.err = arg.args[*expectArgument].SetValueUint(token)
	default:
		arg.err = fmt.Errorf(unsupportedArgumentClass)
	}
}
