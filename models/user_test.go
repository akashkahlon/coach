package models

import (
	"errors"
	"log"
	"testing"
)

func TestUser_Validate(t *testing.T) {
	tests := []struct {
		name     string
		user     User
		expected error
	} { 
		{
			name: "Valid User",
			user: User{
				Name: "John Doe",
				Email: "john@doe.com",
				PasswordHash: "password",
				Role: RegularUser,
			},
			expected: nil,
		},
		{
			name: "Empty Name",
			user: User{
				Name: "",
				Email: "john@doe.com",
				PasswordHash: "password",
				Role: RegularUser,
			},
			expected: errors.New("name is required"),
		},
		{
			name: "Empty Email",
			user: User{
				Name: "John Doe",
				Email: "",
				PasswordHash: "password",
				Role: RegularUser,
			},
			expected: errors.New("email is required"),
		},
		{
			name: "Empty Password",
			user: User{
				Name: "John Doe",
				Email: "john@doe.com",
				PasswordHash: "",
				Role: RegularUser,
			},
			expected: errors.New("password is required"),
		},
		{
			name: "Empty Role",
			user: User{
				Name: "John Doe",
				Email: "john@doe.com",
				PasswordHash: "password",
				Role: "",
			},
			expected: errors.New("role is required"),
		},
		{
			name: "Role is Invalid",
			user: User{
				Name: "John Doe",
				Email: "john@doe.com",
				PasswordHash: "password",
				Role: "superuser",
			},
			expected: errors.New("role is invalid"),
		},
		}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.user.Validate()
			log.Println(err)
			if (err != nil && tt.expected == nil) || (err == nil && tt.expected != nil) || (err != nil && tt.expected != nil && err.Error() != tt.expected.Error()) {
				t.Errorf("Expected %v but got %v", tt.expected, err)
			}
		})
	}
}