package db

import (
	"errors"
	"time"
)

type DbUserDriver struct{}

func NewUserDriver() *DbUserDriver {
	return &DbUserDriver{}
}

type User struct {
	Id        int `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"unique"`
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

func (dbu *DbUserDriver) FindByEmail(email string) (*User, error) {
	var user *User
	// Firstだと存在しない場合にサーバー側でエラーが発生してしまうため、Findでエラーを発生しないようにしている
	result := DB.Where("email = ?", email).Find(&user)
	// 存在しない場合にエラーは発生しないので、エラーを作成する
	if result.RowsAffected == 0 {
		return nil, errors.New("user is not found")
	}
	return user, nil
}

func (dbu *DbUserDriver) UpdateUser(user *User, updateColumns map[string]interface{}) error {
	result := DB.Model(&user).Updates(updateColumns)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (dbu *DbUserDriver) GetUser(id int) (*User, error) {
	var user *User
	// Firstだと存在しない場合にサーバー側でエラーが発生してしまうため、Findでエラーを発生しないようにしている
	result := DB.Find(&user, id)
	// 存在しない場合にエラーは発生しないので、エラーを作成する
	if result.RowsAffected == 0 {
		return nil, errors.New("user is not found")
	}
	return user, nil
}
