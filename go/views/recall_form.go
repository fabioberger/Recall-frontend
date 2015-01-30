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

type RecallForm struct {
	humble.Identifier
}

const (
	reminderFormSelector = "form#reminder-form"
)

func (a *RecallForm) RenderHTML() string {
	return fmt.Sprintf(`
	<form class="form-signin" role="form" id="reminder-form" onsubmit="return false;">
	<div class="alert alert-warning" role="alert" id="form-alert"></div>
  <h2 class="form-signin-heading">What would you like to Recall?</h2>
  <label for="inputEmail" class="sr-only">Knowledge Title</label>
  <input type="text" name="reminder" id="reminder" class="form-control" placeholder="Crypto Course Week 1..." required="" autofocus="">
  <button class="btn btn-lg btn-primary btn-block" type="submit">Recall</button>
</form>
`)
}

func (v *RecallForm) OnLoad() error {
	doc.GetElementByID("form-alert").SetAttribute("style", "display: none;")
	if err := view.AddListener(v, reminderFormSelector, "submit", v.newReminder); err != nil {
		return err
	}
	return nil
}

func (v *RecallForm) newReminder(event dom.Event) {
	event.PreventDefault()
	reminder := doc.GetElementByID("reminder").Underlying().Get("value").String()
	m := &models.Reminder{
		Reminder: reminder,
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
	js.Global.Get("location").Set("hash", "/profile")
}

func (v *RecallForm) OuterTag() string {
	return "div"
}
