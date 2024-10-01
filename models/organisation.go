package models

import (
	"errors"
)

type Organisation struct {
	ID       int
	Name     string
}

func (o *Organisation) Validate() error {
	if o.Name == "" {
		return errors.New("name is required")
	}
	return nil
}