package models

import (
	"strconv"

	"github.com/fabioberger/recall-frontend/go/config"

	"honnef.co/go/js/dom"
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
	root := dom.GetWindow().Location().Hostname
	return "http://" + root + ":" + config.Port + "/sessions"
}
