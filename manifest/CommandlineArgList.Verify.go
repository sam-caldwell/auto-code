package manifest

// Verify - Validate the commandline argument list
func (commands *CommandlineArgList) Verify() error {

	for _, command := range *commands {

		if err := command.Verify(); err != nil {

			return err

		}

	}

	return nil

}
