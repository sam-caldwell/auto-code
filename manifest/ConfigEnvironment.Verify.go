package manifest

// Verify - verify the manifest.yaml definition of an environment variable processor as a data source.
//
// In a manifest.yaml file, environment is configured by something like this--
//
//	environment:
//	 - listenAddress: LISTEN_ADDRESS
//	 - listenPort: LISTEN_PORT
//	 - protocol: LISTEN_PROTOCOL
//
// The list of key-value pairs represents a property name (left-hand side)
// and its environment variable name (right-hand side).
func (environment *ConfigEnvironment) Verify() error {

	for key, value := range *environment {

		if err := key.Verify(); err != nil {
			return err
		}

		if err := value.Verify(); err != nil {
			return err
		}

	}
	return nil
}
