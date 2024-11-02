package manifest

import (
	"github.com/sam-caldwell/auto-code/manifest/validators"
	"testing"
)

func TestCommandlineArgList_Verify(t *testing.T) {

	argMap := CommandlineArgList{
		"foobar": Commandline{
			Required: true,
			Short:    "-f",
			Long:     "--foobar",
		},
	}

	properties := ConfigProperties{
		"foobar": ConfigProperty{
			Type: "string",
			Validator: PropertyValidator{
				Class: "regexp",
				Parameter: validator.Pattern{
					Regex: `[a-z]+`,
					Match: true,
				},
			},
			Default: "testValue",
		},
		"number": ConfigProperty{
			Type: "uint16",
			Validator: PropertyValidator{
				Class: "minmax",
				Parameter: MinMaxUint16{
					Min: 10,
					Max: 20,
				},
			},
			Default: 15,
		},
	}

	if err := argMap.Verify(nil); err != nil {
		t.Fatal(err)
	}

}
