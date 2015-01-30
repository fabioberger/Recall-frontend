package views

import (
	"fmt"

	"honnef.co/go/js/console"
	"honnef.co/go/js/dom"

	"github.com/fabioberger/humble"
	"github.com/fabioberger/humble/model"
	"github.com/fabioberger/humble/view"
	"github.com/fabioberger/recall-frontend/go/models"
	"github.com/gopherjs/gopherjs/js"
)

type Login struct {
	humble.Identifier
}

const (
	logInFormSelector = "form#form-signin"
)

func (a *Login) RenderHTML() string {
	return fmt.Sprintf(`
	<form id="form-signin" onsubmit="return false;">
	<div class="alert alert-warning" role="alert" id="form-alert"></div>
  <h2 class="form-signin-heading">Log In</h2>
  <label for="inputEmail" class="sr-only">Email</label>
  <input type="text" name="email" id="email" class="form-control login-input" placeholder="me@example.com" required="true" autofocus="">
  <input type="password" name="password" id="password" class="form-control login-input" placeholder="password" required="true" autofocus="">
  <div class="checkbox">
    <label>
      <input type="checkbox" name="remember-me" id="remember-me" value="true"> Remember Me
    </label>
  </div>
  <button class="btn btn-lg btn-primary btn-block" type="submit" id="login-submit">Log In</button>
</form>
`)
}

func (v *Login) OnLoad() error {
	doc.GetElementByID("form-alert").SetAttribute("style", "display: none;")
	if err := view.AddListener(v, logInFormSelector, "submit", v.logInClick); err != nil {
		return err
	}
	return nil
}

func (v *Login) logInClick(event dom.Event) {
	console.Log("Logged In Submitted")

	event.PreventDefault()
	email := doc.GetElementByID("email").Underlying().Get("value").String()
	password := doc.GetElementByID("password").Underlying().Get("value").String()
	rememberMe := doc.GetElementByID("remember-me").Underlying().Get("value").Bool()
	m := &models.Session{
		Email:      email,
		Password:   password,
		RememberMe: rememberMe,
	}
	if err := model.Create(m); err != nil {
		addAlert("Could not reach server. Please try again later")
		panic(err)
	}
	if m.Error != "" {
		console.Log(m.Error)
		addAlert(m.Error)
		return
	}
	setUserMenu()
	js.Global.Get("location").Set("hash", "/profile")
}

func addAlert(msg string) {
	doc.GetElementByID("form-alert").SetAttribute("style", "display: block;")
	doc.GetElementByID("form-alert").SetInnerHTML(msg)
}

func (v *Login) OuterTag() string {
	return "div"
}
