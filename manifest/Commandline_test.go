package manifest

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestCommandline_Struct(t *testing.T) {
	t.Run("verify struct fields and types", func(t *testing.T) {

		_ = Commandline{
			Required: true,
			Short:    "",
			Long:     "",
		}

	})
	t.Run("confirm Yaml parse", func(t *testing.T) {

		const testData = "---\n" +
			"required: true\n" +
			"short: shortValue\n" +
			"long: longValue\n"

		var actual Commandline

		if err := yaml.Unmarshal([]byte(testData), &actual); err != nil {
			t.Fatal(err)
		}

		if actual.Required != true {
			t.Fatal("expected true, got false")
		}

		if actual.Short != "shortValue" {
			t.Fatalf("expected shortValue, got %s", actual.Short)
		}

		if actual.Long != "longValue" {
			t.Fatalf("expected longValue, got %s", actual.Long)
		}

	})

}
