package configuration

import (
	"fmt"
	"github.com/sam-caldwell/auto-code/pkg/manifest"
	"github.com/sam-caldwell/auto-code/pkg/manifest/pdo"
	"gopkg.in/yaml.v3"
	"strings"
)

// UnmarshalYAML - unmarshal the YAML Parameter Value block.
func (p *ParameterValue) UnmarshalYAML(node *yaml.Node) error {

	/*
	 * helper functions
	 */

	trimElements := func(listOfElements []string) []string {
		for i, v := range listOfElements {
			listOfElements[i] = strings.TrimSpace(v)
		}
		return listOfElements
	}

	parseEnumValue := func() error {
		// Capture and parse any string 'enum(element1,element2,...elementN)
		const (
			prefix    = "enum("
			suffix    = ")"
			delimiter = ","
		)
		if strings.HasPrefix(p.Data.String(), prefix) && strings.HasSuffix(p.Data.String(), suffix) {
			parts := trimElements(strings.Split(
				strings.TrimSuffix(
					strings.TrimPrefix(
						strings.TrimSpace(p.Data.String()), prefix),
					suffix),
				delimiter))

			for i, _ := range parts {
				parts[i] = strings.TrimSpace(parts[i])
			}
			p.Data.State = make(pdo.Enum, len(parts))
		}
		return nil
	}

	parseArrayValue := func() (err error) {
		switch p.Data.State.(type) {
		case []any, pdo.Array:
			p.Data.State = pdo.Array(p.Data.State.([]any))
		default:
			err = fmt.Errorf("invalid input (expected array)")
		}
		return err
	}

	parseObjectValue := func() (err error) {
		switch p.Data.State.(type) {
		case struct{}:
			p.Data.State = p.Data.State.(pdo.PdoParameterObject)
		default:
			err = fmt.Errorf("invalid input (expected object)")
		}
		return err
	}

	/*
	 * function payload
	 */

	// Parse the YAML...
	if err := manifest.ParseYamlObjectWithReferences(node, p); err != nil {
		return err
	}

	// Fix up the special (Array, enum, object) types
	switch p.Data.State.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, bool, float32, float64:
		return nil
	case string:
		return parseEnumValue()
	case []any:
		return parseArrayValue()
	case struct{}:
		return parseObjectValue()
	default:
		panic("unknown type")
	}
}
