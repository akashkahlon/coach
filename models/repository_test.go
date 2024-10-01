package models

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestRepository_Validate(t *testing.T) {
	tests := []struct {
		name       string
		repository Repository
		expected   error
	}{
		{
			name: "Valid Repository",
			repository: Repository{
				Name: "John Doe",
				FullName: "John Doe",
				OrganisationID: 1,
				Private: false,
				HtmlURL: "https://github.com/shurutech/primus",
				GithubDataAll: json.RawMessage(`{"key": "value", "key2": "value2"}`),
			},
			expected: nil,
		},
		{
			name: "Empty Name",
			repository: Repository{
				Name: "",
				FullName: "John Doe",
				OrganisationID: 1,
				Private: false,
				HtmlURL: "https://github.com/shurutech/primus",
				GithubDataAll: json.RawMessage(`{"key": "value", "key2": "value2"}`),
			},
			expected: errors.New("name is required"),
		},
		{
			name: "Empty FullName",
			repository: Repository{
				Name: "John Doe",
				FullName: "",
				OrganisationID: 1,
				Private: false,
				HtmlURL: "https://github.com/shurutech/primus",
				GithubDataAll: json.RawMessage(`{"key": "value", "key2": "value2"}`),
		},
			expected: errors.New("full name is required"),
		},
		{
			name: "Empty OrganisationID",
			repository: Repository{
				Name: "John Doe",
				FullName: "John Doe",
				OrganisationID: 0,
				Private: false,
				HtmlURL: "https://github.com/shurutech/primus",
				GithubDataAll: json.RawMessage(`{"key": "value", "key2": "value2"}`),
		},
			expected: errors.New("organisation id is required"),
		},
		{
			name: "Empty HtmlURL",
			repository: Repository{
				Name: "John Doe",
				FullName: "John Doe",
				OrganisationID: 1,
				Private: false,
				HtmlURL: "",
				GithubDataAll: json.RawMessage(`{"key": "value", "key2": "value2"}`),
		},
			expected: errors.New("html url is required"),
		},
		{
			name: "Empty GithubDataAll",
			repository: Repository{
				Name: "John Doe",
				FullName: "John Doe",
				OrganisationID: 1,
				Private: false,
				HtmlURL: "https://github.com/shurutech/primus",
				GithubDataAll: nil,
		},
			expected: errors.New("github data all is required"),
		},
		{
			name: "Empty GithubDataAll",
			repository: Repository{
				Name: "John Doe",
				FullName: "John Doe",
				OrganisationID: 1,
				Private: false,
				HtmlURL: "https://github.com/shurutech/primus",
				GithubDataAll: json.RawMessage([]byte{}),
		},
			expected: errors.New("github data all is required"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.repository.Validate()
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