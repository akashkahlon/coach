package models

import (
	"encoding/json"
	"errors"
	"time"
)

type Repository struct {
	ID             int            `gorm:"primaryKey"`
	Name           string         `gorm:"not null"`
	FullName       string         `gorm:"not null"`
	OrganisationID int            `gorm:"not null"`
	Private        bool           `gorm:"not null"`
	HtmlURL        string         `gorm:"not null"`
	GithubDataAll  json.RawMessage `gorm:"type:jsonb"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (a *Repository) Validate() error {
	if a.Name == "" {
		return errors.New("name is required")
	}
	if a.FullName == "" {
		return errors.New("full name is required")
	}
	if a.OrganisationID == 0 {
		return errors.New("organisation id is required")
	}
	if a.HtmlURL == "" {
		return errors.New("html url is required")
	}
	if len(a.GithubDataAll) == 0 || a.GithubDataAll == nil {
		return errors.New("github data all is required")
	}

	return nil
}
