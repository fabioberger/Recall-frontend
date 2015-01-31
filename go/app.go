package main

import (
	"github.com/fabioberger/humble/model"
	"github.com/fabioberger/humble/router"
	"github.com/fabioberger/humble/view"
	"github.com/fabioberger/recall-frontend/go/models"
	"github.com/fabioberger/recall-frontend/go/views"
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/console"
	"honnef.co/go/js/dom"
)

const (
	bodySelector = "body"
)

var (
	doc      = dom.GetWindow().Document()
	elements = struct {
		body dom.Element
	}{}
	appHasLoaded = false
)

// Get the body element
func init() {
	elements.body = doc.QuerySelector(bodySelector)
}

func main() {
	console.Log("Starting...")

	//Start main app view, appView
	appView := &views.App{}
	if err := view.ReplaceParentHTML(appView, bodySelector); err != nil {
		panic(err)
	}

	// Routes
	r := router.New()
	r.HandleFunc("/", func(params map[string]string) {
		form := &views.RecallForm{}
		view.ReplaceParentHTML(form, "#container")
	})
	r.HandleFunc("/login", func(params map[string]string) {
		// Set up the footer view
		login := &views.Login{}
		view.ReplaceParentHTML(login, "#container")
	})
	r.HandleFunc("/signup", func(params map[string]string) {
		// Set up the footer view
		signup := &views.Signup{}
		view.ReplaceParentHTML(signup, "#container")
	})
	r.HandleFunc("/profile", func(params map[string]string) {
		user := &models.User{}
		if err := model.Read(user, 0); err != nil {
			panic(err)
		}
		if user.Error != "" {
			console.Log("User doens't exist")
			js.Global.Get("location").Set("hash", "/")
			return
		}
		reminders := []*models.Reminder{}
		if err := model.ReadAll(&reminders); err != nil {
			panic(err)
		}
		console.Log(reminders)
		profile := &views.Profile{
			User:      user,
			Reminders: reminders,
		}
		profile.InitChildren(reminders)
		view.ReplaceParentHTML(profile, "#container")
	})
	r.Start()

}
