package model

import (
	"errors"
	"regexp"
)

type User struct {
	Id    int
	Name  string
	Email string
}

func (u User) Valid() error {
    emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
    re := regexp.MustCompile(emailRegex)
    if !re.MatchString(u.Email) {
        return errors.New("emailではありません")
    }
	return nil
}

func NewUser(name string, email string) (*User, error) {
	user := &User{
		Name:  name,
		Email: email,
	}
	if err := user.Valid(); err != nil {
		return &User{}, err
	}
	return user, nil
}
