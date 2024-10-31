package arguments

// storeValue - Store a token as a value if it is not an argument since we expect a value.
func (config *Configuration) storeValue(arg *Argument, token *string, expectValue *bool, expectArgument *string, currentClass *ArgumentClass) (err error) {
	//ToDo: totally rewrite this.
	//if detectArgument(token) {
	//	return fmt.Errorf(unknownOrUnexpectedArgument)
	//}
	//switch currentArgument := config.data[*expectArgument]; currentArgument.class {
	//case Bool:
	//	arg.err = arg.args[*expectArgument].SetValueBool(token)
	//case Flag:
	//	trueString := "true"
	//	arg.err = arg.args[*expectArgument].SetValueBool(&trueString)
	//case Float:
	//	arg.err = arg.args[*expectArgument].SetValueFloat(token)
	//case Int:
	//	if value, err := strconv.ParseInt(*token, 10, 64); err != nil {
	//		arg.err = err
	//	} else {
	//		err = arg.args[*expectArgument].SetValueInt(value)
	//	}
	//case String, Directory, File, Secret, Email:
	//	err = arg.args[*expectArgument].SetValueString(token)
	//case Uint:
	//	err = arg.args[*expectArgument].SetValueUint(token)
	//default:
	//	err = fmt.Errorf(unsupportedArgumentClass)
	//}
	return err
}
