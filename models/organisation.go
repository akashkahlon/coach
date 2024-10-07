package models

import (
	"errors"
)

type Organisation struct {
	ID       int      `gorm:"primaryKey"`
	Name     string   `gorm:"not null"`
}

func (o *Organisation) Validate() error {
	if o.Name == "" {
		return errors.New("name is required")
	}
	return nil
}