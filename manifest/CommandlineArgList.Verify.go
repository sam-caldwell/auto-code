package manifest

// Verify - Validate the commandline argument list
func (commands *CommandlineArgList) Verify(properties *ConfigProperties) error {

	for _, command := range *commands {

		if err := command.Verify(properties); err != nil {

			return err

		}

	}

	return nil

}
