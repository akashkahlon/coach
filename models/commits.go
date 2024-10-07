package models

import (
	"encoding/json"
	"errors"
	"time"
)
type Commit struct {
	ID            int            `gorm:"primaryKey"`
	RepositoryID  int            `gorm:"not null"`
	Sha           string         `gorm:"not null"`
	Author        string         `gorm:"not null"`
	AuthorEmail   string         `gorm:"not null"`
	CommitMessage string         `gorm:"not null"`
	GithubDataAll json.RawMessage `gorm:"type:jsonb"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
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
