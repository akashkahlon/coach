package models

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestCommit_Validate(t *testing.T) {
	tests := []struct {
		name     string
		commit   Commit
		expected error
	}{
		{
			name: "Valid Commit",
			commit: Commit{
				RepositoryID: 1,
				Sha: 				"abc123",
				Author: 		"John Doe",
				AuthorEmail: "john@doe.com",
				CommitMessage:        "Initial commit",
				GithubDataAll: json.RawMessage(`{"key": "value", "key2": "value2"}`),
			},
			expected: nil,
		},
		{
			name: "Empty RepositoryID",
			commit: Commit{
				RepositoryID: 0,
				Sha: 				"abc123",
				Author: 		"John Doe",
				AuthorEmail: "john@doe.com",
				CommitMessage:        "Initial commit",
				GithubDataAll: json.RawMessage(`{"key": "value", "key2": "value2"}`),
	},
			expected: errors.New("repository id is required"),
		},
		{
			name: "Empty Sha",
			commit: Commit{
				RepositoryID: 1,
				Sha: 				"",
				Author: 		"John Doe",
				AuthorEmail: "john@doe.com",
				CommitMessage:        "Initial commit",
				GithubDataAll: json.RawMessage(`{"key": "value", "key2": "value2"}`),
		},
			expected: errors.New("sha is required"),
		},
		{
			name: "Empty Author",
			commit: Commit{
				RepositoryID: 1,
				Sha: 				"abc123",
				Author: 		"",
				AuthorEmail: "joh@doe.com",
				CommitMessage:        "Initial commit",
				GithubDataAll: json.RawMessage(`{"key": "value", "key2": "value2"}`),
		},
			expected: errors.New("author is required"),
		},
		{
			name: "Empty AuthorEmail",
			commit: Commit{
				RepositoryID: 1,
				Sha: 				"abc123",
				Author: 		"John Doe",
				AuthorEmail: "",
				CommitMessage:        "Initial commit",
				GithubDataAll: json.RawMessage(`{"key": "value", "key2": "value2"}`),

		},
			expected: errors.New("author email is required"),
		},
		{
			name: "Empty CommitMessage",
			commit: Commit{
				RepositoryID: 1,
				Sha: 				"abc123",
				Author: 		"John Doe",
				AuthorEmail: "john@doe.com",
				CommitMessage:        "",
				GithubDataAll: json.RawMessage(`{"key": "value", "key2": "value2"}`),
		},
			expected: errors.New("commit message is required"),
		},
		{
			name: "Missing GithubDataAll",
			commit: Commit{
				RepositoryID: 1,
				Sha: 				"abc123",
				Author: 		"John Doe",
				AuthorEmail: "john@doe.com",
				CommitMessage:        "Initial commit",
				GithubDataAll: nil,
		},
			expected: errors.New("github data all is required"),
		},
		{
			name: "Empty GithubDataAll",
			commit: Commit{
				RepositoryID: 1,
				Sha: 				"abc123",
				Author: 		"John Doe",
				AuthorEmail: "john@doe.com",
				CommitMessage:        "Initial commit",
				GithubDataAll: json.RawMessage([]byte{}),
		},
			expected: errors.New("github data all is required"),
	},
	}


	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.commit.Validate()
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