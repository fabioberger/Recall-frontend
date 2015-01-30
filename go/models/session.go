package models

import (
	"strconv"
)

type Session struct {
	Id         int
	Email      string
	Password   string
	RememberMe bool
	Error      string
}

func (t *Session) GetId() string {
	return strconv.Itoa(t.Id)
}

func (t *Session) RootURL() string {
	return "http://localhost:8080/sessions"
}
