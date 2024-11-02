package manifest

import (
	"fmt"
	"strings"
)

// Verify - Validate the command line properties definition.
func (command *Commandline) Verify(_ *ConfigProperties) error {

	command.Short = ShortArgumentString(strings.TrimSpace(string(command.Short)))

	command.Long = LongArgumentString(strings.TrimSpace(string(command.Long)))

	if (command.Short) == EmptyString && string(command.Long) == EmptyString {

		return fmt.Errorf(errMissingCommandlineArgument)

	}

	if err := command.Short.Verify(); err != nil {

		return err

	}

	if err := command.Long.Verify(); err != nil {

		return err

	}

	return nil
}
