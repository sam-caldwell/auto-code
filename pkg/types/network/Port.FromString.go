package network

import "strconv"

// FromString - convert a string port number to its numeric state
func (n *Port) FromString(portString string) error {

	if parsedPort, err := strconv.Atoi(portString); err == nil {

		*n = Port(parsedPort)
		return nil

	} else {

		return err

	}

}
