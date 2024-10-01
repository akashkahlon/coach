package models

import (
	"encoding/json"
	"errors"
)

type Commit struct {
	ID       int
	RepositoryID int
	Sha string
	Author string
	AuthorEmail string
	CommitMessage string
	GithubDataAll json.RawMessage
}

func (c *Commit) Validate() error {
	if c.RepositoryID == 0 {
		return errors.New("repository id is required")
	}
	if c.Sha == "" {
		return errors.New("sha is required")
	}
	if c.Author == "" {
		return errors.New("author is required")
	}
	if c.AuthorEmail == "" {
		return errors.New("author email is required")
	}
	if c.CommitMessage == "" {
		return errors.New("commit message is required")
	}
	if len(c.GithubDataAll) == 0 || c.GithubDataAll == nil {
		return errors.New("github data all is required")
	}

	return nil
}
