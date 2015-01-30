package views

import (
	"fmt"

	"honnef.co/go/js/dom"

	"github.com/fabioberger/humble"
	"github.com/fabioberger/humble/model"
	"github.com/fabioberger/humble/view"
	"github.com/fabioberger/recall-frontend/go/models"
	"github.com/gopherjs/gopherjs/js"
)

type Signup struct {
	humble.Identifier
}

const (
	newUserSelector = "form#form-signup"
)

func (a *Signup) RenderHTML() string {
	return fmt.Sprintf(`
		<form class="form-signin" role="form" id="form-signup" onsubmit="return false;">
	  <h2 class="form-signin-heading">Sign Up</h2>
	  <div class="alert alert-warning" role="alert" id="form-alert"></div>
	  <label for="inputName" class="sr-only">Name</label>
	  <input type="text" name="name" id="name" class="form-control" placeholder="Bob Jones" required="true" autofocus="">
	  <label for="inputEmail" class="sr-only">Email</label>
	  <input type="text" name="email" id="email" class="form-control" placeholder="me@example.com" required="true" autofocus="">
	  <input type="password" name="password" id="password" class="form-control" placeholder="password" required="true" autofocus="">
	   <input type="password" name="confirmPassword" id="confirm_password" class="form-control" placeholder="confirm password" required="true" autofocus="">
	  <button class="btn btn-lg btn-primary btn-block" id="new-user" type="submit">Log In</button>
	  </form>
`)
}

func (v *Signup) OuterTag() string {
	return "div"
}

func (v *Signup) OnLoad() error {
	doc.GetElementByID("form-alert").SetAttribute("style", "display: none;")
	// Add listener on signup submission
	if err := view.AddListener(v, newUserSelector, "submit", v.newUserClick); err != nil {
		return err
	}
	return nil
}

func (v *Signup) newUserClick(event dom.Event) {
	event.PreventDefault()
	name := doc.GetElementByID("name").Underlying().Get("value").String()
	email := doc.GetElementByID("email").Underlying().Get("value").String()
	password := doc.GetElementByID("password").Underlying().Get("value").String()
	confirmPassword := doc.GetElementByID("confirm_password").Underlying().Get("value").String()
	//Create a model, send to server and append view
	m := &models.User{
		Name:            name,
		Email:           email,
		Password:        password,
		ConfirmPassword: confirmPassword,
	}
	if err := model.Create(m); err != nil {
		addAlert("Could not reach server. Please try again later")
		panic(err)
	}
	if m.Error != "" {
		addAlert(m.Error)
		return
	}
	setUserMenu()
	js.Global.Get("location").Set("hash", "/profile")
}

func setUserMenu() {
	doc.GetElementByID("loggedOut").SetAttribute("style", "display: none;")
	doc.GetElementByID("loggedIn").SetAttribute("style", "display: block;")
}

func setVisitorMenu() {
	doc.GetElementByID("loggedOut").SetAttribute("style", "display: block;")
	doc.GetElementByID("loggedIn").SetAttribute("style", "display: none;")
}
