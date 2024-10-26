package arguments

import (
	"fmt"
	"strconv"
	"strings"
)

func (arg *ArgumentDescriptor[I, U, F, S, B]) SetValueBool(raw *string) error {
	switch v := strings.ToLower(strings.TrimSpace(*raw)); v {
	case "true", "false":
		arg.value = *raw
		return nil
	default:
		return fmt.Errorf(invalidBooleanValue)
	}
}

func (arg *ArgumentDescriptor[I, U, F, S, B]) SetValueFloat(raw *string) error {
	if v, err := strconv.ParseFloat(strings.TrimSpace(*raw), 32); err != nil {
		return err
	} else {
		arg.value = v
	}
	if arg.bounds.Validator(arg.value) {
		return nil
	}
	return fmt.Errorf(boundsCheckFailed, arg.flag, arg.value)
}

func (arg *ArgumentDescriptor[I, U, F, S, B]) SetValueInt(raw *string) error {
	if v, err := strconv.ParseInt(strings.TrimSpace(*raw), 10, 64); err != nil {
		return err
	} else {
		arg.value = v
	}
	if arg.bounds.Validator(arg.value) {
		return nil
	}
	return fmt.Errorf(boundsCheckFailed, arg.flag, arg.value)
}

func (arg *ArgumentDescriptor[I, U, F, S, B]) SetValueString(raw *string) error {
	if err := arg.Validator(); err != nil {
		return err
	}
	arg.value = *raw
	return nil
}

func (arg *ArgumentDescriptor[I, U, F, S, B]) SetValueUint(raw *string) error {
	if v, err := strconv.ParseUint(strings.TrimSpace(*raw), 10, 64); err != nil {
		return err
	} else {
		arg.value = v
	}
	if arg.bounds.Validator(arg.value) {
		return nil
	}
	return fmt.Errorf(boundsCheckFailed, arg.flag, arg.value)
}
