package models

import (
	"errors"
	"fmt"
	"html"
	"strings"

	"github.com/jinzhu/gorm"
	token "github.com/rest-api/utils"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Username     string `gorm:"size:255;not null;unique" json:"username"`
	Email        string `gorm:"size:255;not null;unique" json:"email"`
	Mobilenumber string `gorm:"size:255;not null;unique" json:"mobilnumber"`
	Password     string `gorm:"size:255;not null;" json:"password"`
}

func GetUserByID(uid uint) (User, error) {

	var u User

	if err := DB.First(&u, uid).Error; err != nil {
		return u, errors.New("User not found!")
	}

	u.PrepareGive()

	return u, nil

}

func (u *User) PrepareGive() {
	u.Password = ""
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username string, password string) (string, error) {

	var err error

	u := User{}

	err = DB.Model(User{}).Where("username = ?", username).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(u.ID)

	fmt.Println(err)

	if err != nil {
		return "", err
	}

	return token, nil

}

func (u *User) SaveUser() (*User, error) {

	var err error
	err = DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) UpdateUserData() (*User, error) {

	fmt.Println(u)

	err := DB.Model(&User{}).Updates(u).Error

	if err != nil {
		return &User{}, err
	}

	return &User{}, nil
}

func (u *User) DeleteUserD() error {

	err := DB.Model(&User{}).Delete(u).Error

	if err != nil {
		return err
	}

	return nil
}

func (u *User) BeforeSave() error {

	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	//remove spaces in username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil

}
