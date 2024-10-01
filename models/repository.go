package models

import (
	"encoding/json"
	"errors"
)

type Repository struct {
	ID       int
	Name     string
	FullName string
	OrganisationID int
	Private	bool
	HtmlURL string
	GithubDataAll json.RawMessage
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
