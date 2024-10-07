package models

import (
	"errors"
	"time"
)
type Account struct {
	ID       int `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (a *Account) Validate() error {
	if a.Name == "" {
		return errors.New("name is required")
	}
	return nil
}