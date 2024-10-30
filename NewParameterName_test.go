package arguments

import (
	"fmt"
	"github.com/google/uuid"
	"testing"
)

func TestNewParameterName(t *testing.T) {

	t.Run("ensure UUIDs pass validation since we use UUIDs for temporary parameterNames", func(t *testing.T) {

		id, _ := uuid.NewUUID()

		if _, err := NewParameterName(id.String()); err != nil {
			t.Fatalf("Expected no error, but got: %v", err)
		}
	})

	t.Run("Happy Path: valid ParameterName", func(t *testing.T) {

		const validName = "valid_parameter-name.1"
		param, err := NewParameterName(validName)

		if err != nil {
			t.Fatalf("Expected no error for valid parameter name, got: %v", err)
		}

		if param == nil {
			t.Fatalf("Expected parameter to be non-nil")
		}

		if string(*param) != validName {
			t.Fatalf("Expected parameter name to match the input; got: %s, want: %s", *param, validName)
		}
	})

	t.Run("Sad Path: invalid parameter names", func(t *testing.T) {

		var invalidNames []string
		t.Run("Setup test(sad path): create an invalidNames list", func(t *testing.T) {

			// Create a list of invalid names using invalid characters.
			invalidCharacters := []rune{
				' ', '!', '"', '#', '$', '%', '&', '\'', '(', ')', '*', '+', ',', '/', ':',
				';', '<', '=', '>', '?', '@', '[', '\\', ']', '^', '`', '{', '|', '}', '~',
			}

			for _, invalidChar := range invalidCharacters {
				invalidNames = append(invalidNames,
					fmt.Sprintf("%c_leading.bad-char5", invalidChar),
					fmt.Sprintf("middle_b.%c.ad1-char", invalidChar),
					fmt.Sprintf("trailing_bad1-char3.%c", invalidChar),
				)
			}
		})

		t.Run("sad path: Verify that the invalidNames result in errors", func(t *testing.T) {

			// iterate over the invalidNames list and test each.
			for _, invalidName := range invalidNames {

				t.Run("sad path: test with "+invalidName, func(t *testing.T) {

					isValidUUID := func(str string) bool {
						_, err := uuid.Parse(str)
						return err == nil
					}

					param, err := NewParameterName(invalidName)

					if err == nil {
						t.Fatalf("Expected an error for invalid parameter name (%s), but got none", invalidName)
					}

					if param == nil {
						t.Fatalf("Expected parameter to be non-nil for '%s'", invalidName)
					}

					if !isValidUUID(string(*param)) {
						t.Fatalf("Expected UUID as temporary parameter name; got: %s on %s", *param, invalidName)
					}
				})
			}
		})
	})
}
