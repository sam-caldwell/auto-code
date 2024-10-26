package arguments

import "fmt"

func (arg *Arguments) storePositionalArgument(token *string, position uint, expectValue *bool, expectArgument *string, currentClass *ArgumentClass) {
	// ToDo: check if any positional argument exists at this location.

	//var actualArgument ArgumentDescriptor[int64, uint64, float64, string, bool]
	//if actualArgument, ok := arg.args[token]; ok {
	//arg.err = fmt.Errorf(unknownOrUnexpectedArgument)
	//break // this is an error state

	// Default to an error until we implement positional arguments.
	arg.err = fmt.Errorf(unknownOrUnexpectedArgument)
	return
}
