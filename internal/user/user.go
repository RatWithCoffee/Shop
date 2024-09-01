package user

import (
	"net/mail"
)

type User struct {
	ID       uint   `json:"-" gorm:"primarykey"`
	Email    string `json:"email" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
	Username string `json:"username" gorm:"not null"`
}

func (user User) GetId() int {
	return int(user.ID)
}

func isValidUser(user User) bool {
	addr, ok := validMailAddress(user.Email)
	if ok {
		user.Email = addr
	}
	if !ok || len(user.Username) == 0 || len(user.Password) == 0 {
		return false
	}

	return true
}

func validMailAddress(address string) (string, bool) {
	addr, err := mail.ParseAddress(address)
	if err != nil {
		return "", false
	}
	return addr.Address, true
}
