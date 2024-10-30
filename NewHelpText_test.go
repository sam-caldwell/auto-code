package arguments

import (
	"testing"
)

func TestNewHelpText(t *testing.T) {
	t.Run("Happy Path: valid help text", func(t *testing.T) {
		validHelpText := "This parameter specifies the file path."
		helpTextPtr := &validHelpText

		ht, err := NewHelpText(helpTextPtr)

		if err != nil {
			t.Fatalf("Expected no error for valid help text, got: %v", err)
		}
		if ht == nil {
			t.Fatalf("Expected HelpText to be non-nil")
		}
		if string(*ht) != validHelpText {
			t.Fatalf("Expected HelpText to match input; got: %s, want: %s", *ht, validHelpText)
		}
	})

	t.Run("Sad Path: empty or whitespace-only help text", func(t *testing.T) {
		invalidHelpTexts := []string{
			"",         // Empty string
			"   ",      // Only spaces
			"\t\n  \t", // Whitespace characters
		}

		for _, invalidText := range invalidHelpTexts {
			t.Run("invalid help text: "+invalidText, func(t *testing.T) {
				helpTextPtr := &invalidText

				ht, err := NewHelpText(helpTextPtr)

				if err == nil {
					t.Fatalf("Expected an error for invalid help text, but got none")
				}
				if ht == nil {
					t.Fatalf("Expected HelpText to be non-nil")
				}
				if string(*ht) != "" {
					t.Fatalf("Expected empty HelpText as temporary value; got: %s", *ht)
				}
			})
		}
	})
}
