package models

import (
	"strconv"

	"github.com/fabioberger/recall-frontend/go/config"
	"honnef.co/go/js/dom"
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
	root := dom.GetWindow().Location().Hostname
	return "http://" + root + ":" + config.Port + "/users"
}
