package dataDriver

import (
	"errors"
	"gopkg.in/yaml.v3"
	"strings"
)

// UnmarshalYAML - Unmarshal a yaml node to this DataDriver object
func (d *DataDriver) UnmarshalYAML(node *yaml.Node) error {
	var value string
	if err := node.Decode(&value); err != nil {
		return err
	}
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "file":
		*d = File
	case "postgres":
		*d = Postgres
	case "aws-s3":
		*d = AwsS3
	default:
		return errors.New("unknown dataContract driver: " + value)
	}
	return nil
}
