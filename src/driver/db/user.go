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
	Age       int
	Sex       float32
	Gender    float32
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (dbu *DbUserDriver) CreateUser(user *User) (*User, error) {
	result := DB.Create(&user)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return user, err
}

func (dbu *DbUserDriver) FindByEmail(email string) error {
	var user []*User
	// 一致するemailがあるかを確認する
	result := DB.Where("email = ?", email).First(&user)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}
