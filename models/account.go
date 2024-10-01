package models

import (
	"errors"
)

type Account struct {
	ID       int
	Name     string
}

func (a *Account) Validate() error {
	if a.Name == "" {
		return errors.New("name is required")
	}
	return nil
}