package models

import (
	"strconv"
)

type User struct {
	Id              int
	Name            string
	Email           string
	Password        string
	ConfirmPassword string
	Error           string
}

func (t *User) GetId() string {
	return strconv.Itoa(t.Id)
}

func (t *User) RootURL() string {
	return "http://localhost:8080/users"
}
