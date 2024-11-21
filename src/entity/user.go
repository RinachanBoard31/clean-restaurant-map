package model

import (
	"errors"
	"regexp"
	"strconv"
)

type User struct {
	Id     int
	Name   string
	Email  string
	Age    int     // xx代として表記する(60代以上は全て60とする)
	Sex    float32 // -1.0(男性)~1.0(女性)で表現する。中性、無回答は0となる。
	Gender float32 // -1.0(男性)~1.0(女性)で表現する。中性、無回答は0となる。
}

type UserCredentials struct {
	Email string
}

func emailValid(email string) error {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(email) {
		return errors.New("emailではありません")
	}
	return nil
}

func ageValid(age int) error {
	if age < 0 {
		return errors.New("年齢が0未満です。")
	}
	return nil
}

func ageFormat(age int) int {
	// ageValidで0未満はエラーとなるので0未満は扱わない。
	if age >= 60 {
		return 60
	}
	return (age / 10) * 10
}

func sexFormat(sex float32) float32 {
	if sex < -1.0 {
		return -1.0
	}
	if sex > 1.0 {
		return 1.0
	}
	return sex
}

func genderFormat(gender float32) float32 {
	if gender < -1.0 {
		return -1.0
	}
	if gender > 1.0 {
		return 1.0
	}
	return gender
}

func NewUser(name string, email string, age int, sex float32, gender float32) (*User, error) {
	// バリデーションのチェック
	emailValidError := emailValid(email)
	ageValidError := ageValid(age)
	if err := errors.Join(emailValidError, ageValidError); err != nil {
		return nil, err
	}
	// userの作成
	user := &User{
		Name:   name,
		Email:  email,
		Age:    ageFormat(age),
		Sex:    sexFormat(sex),
		Gender: genderFormat(gender),
	}
	return user, nil
}

func NewUserCredentials(email string) (*UserCredentials, error) {
	// バリデーションのチェック
	emailValidError := emailValid(email)
	if err := errors.Join(emailValidError); err != nil {
		return &UserCredentials{}, err
	}
	user := &UserCredentials{
		Email: email,
	}
	return user, nil
}

func UserFormat(data map[string]interface{}) (map[string]interface{}, error) {
	formatedData := map[string]interface{}{}
	// フォーマットの変換
	// id
	if id, ok := data["id"].(string); ok {
		formatedData["id"], _ = strconv.Atoi(id)
	} else if id, ok := data["id"].(float64); ok {
		formatedData["id"] = int(id)
	} else {
		formatedData["id"] = data["id"]
	}
	// name
	if name, ok := data["name"].(float64); ok {
		formatedData["name"] = strconv.FormatFloat(name, 'f', -1, 64)
	} else {
		formatedData["name"] = data["name"]
	}
	// email
	if email, ok := data["email"]; ok {
		formatedData["email"] = email
	}
	// age
	if age, ok := data["age"].(string); ok {
		formatedData["age"], _ = strconv.Atoi(age)
	} else if age, ok := data["age"].(float64); ok {
		formatedData["age"] = int(age)
	} else {
		formatedData["age"] = data["age"]
	}
	// sex
	if sex, ok := data["sex"].(string); ok {
		formatedData["sex"], _ = strconv.ParseFloat(sex, 32)
	} else if sex, ok := data["sex"].(float64); ok {
		formatedData["sex"] = sexFormat(float32(sex))
	} else {
		formatedData["sex"] = data["sex"]
	}
	// gender
	if gender, ok := data["gender"].(string); ok {
		formatedData["gender"], _ = strconv.ParseFloat(gender, 32)
	} else if gender, ok := data["gender"].(float64); ok {
		formatedData["gender"] = genderFormat(float32(gender))
	} else {
		formatedData["gender"] = data["gender"]
	}

	// バリデーションのチェック
	var emailValidError error = nil
	var ageValidError error = nil
	if email, ok := formatedData["email"].(string); ok {
		emailValidError = emailValid(email)
	}
	if age, ok := formatedData["age"].(int); ok {
		ageValidError = ageValid(age)
	}
	if err := errors.Join(emailValidError, ageValidError); err != nil {
		return nil, err
	}

	return formatedData, nil
}

func UpdateConditions(data map[string]interface{}) error {
	// idが含まれている
	// emailは更新不可
	if _, ok := data["id"]; !ok {
		return errors.New("id is not found")
	}
	if _, ok := data["email"]; ok {
		delete(data, "email")
	}
	return nil
}
