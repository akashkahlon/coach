package models

import (
	"errors"
	"time"
)

type UserRole string

const (
	Admin UserRole = "admin"
	RegularUser  UserRole = "user"
)

type User struct {
	ID           int       `gorm:"primaryKey"`
	Name         string    `gorm:"not null"`
	Email        string    `gorm:"not null;unique"`
	PasswordHash string    `gorm:"not null"`
	Role         UserRole  `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (u *User) Validate() error {
	if u.Name == "" {
		return errors.New("name is required")
	}
	if u.Email == "" {
		return errors.New("email is required")
	}
	if u.PasswordHash == "" {
		return errors.New("password is required")
	}
	if u.Role == "" {
		return errors.New("role is required")
	}
	if u.Role != Admin && u.Role != RegularUser {
		return errors.New("role is invalid")
	}
	return nil
}
