package db

import (
	"time"
)

type DbUserDriver struct{}

func NewUserDriver() *DbUserDriver {
	return &DbUserDriver{}
}

type User struct {
	Id        int `gorm:"primaryKey"`
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (dbu *DbUserDriver) CreateUser(user *User) error {
	err := DB.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}
