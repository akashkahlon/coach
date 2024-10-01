package models

import (
	"errors"
	"testing"
)

func TestAccount_Validate(t *testing.T) {
	tests := []struct {
		name     string
		account  Account
		expected error
	}{
		{
			name: "Valid Account",
			account: Account{
				Name: "John Doe",
			},
			expected: nil,
		},
		{
			name: "Empty Name",
			account: Account{
				Name: "",
			},
			expected: errors.New("name is required"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.account.Validate()
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