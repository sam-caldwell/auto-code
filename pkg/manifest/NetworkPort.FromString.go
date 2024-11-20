package manifest

import "strconv"

// FromString - convert a string port number to its numeric state
func (n *NetworkPort) FromString(portString string) error {
	if parsedPort, err := strconv.Atoi(portString); err == nil {
		*n = NetworkPort(parsedPort)
		return nil
	} else {
		return err
	}
}
