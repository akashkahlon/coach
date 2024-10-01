package models

import (
	"errors"
	"testing"
)

func TestOrganisation_Validate(t *testing.T) {
	tests := []struct {
		name     string
		organisation Organisation
		expected error
	}{
		{
			name: "Valid Organisation",
			organisation: Organisation{
				Name: "John Doe",
			},
			expected: nil,
		},
		{
			name: "Empty Name",
			organisation: Organisation{
				Name: "",
			},
			expected: errors.New("name is required"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.organisation.Validate()
			if tt.expected != nil {
				if err.Error() != tt.expected.Error() {
					t.Errorf("Expected %v but got %v", tt.expected, err)
				}
			} else if err != nil {
				t.Errorf("Expected no error but got %v", err)
			}
		})
	}
}