package manifest

import (
	"fmt"
	data2 "github.com/sam-caldwell/auto-code/data"
	"github.com/sam-caldwell/auto-code/messages"
	"github.com/sam-caldwell/auto-code/words"
	"strings"
)

// Verify - Validate the command line properties definition.
func (command *Commandline) Verify(_ *ConfigProperties) error {

	command.Short = data2.ShortArgumentString(strings.TrimSpace(string(command.Short)))

	command.Long = data2.LongArgumentString(strings.TrimSpace(string(command.Long)))

	if (command.Short) == words.EmptyString && string(command.Long) == words.EmptyString {

		return fmt.Errorf(messages.ErrMissingCommandlineArgument)

	}

	if err := command.Short.Verify(); err != nil {

		return err

	}

	if err := command.Long.Verify(); err != nil {

		return err

	}

	return nil
}
