package views

import (
	"fmt"
	"strings"

	"github.com/fabioberger/humble"
	"github.com/fabioberger/humble/model"
	"github.com/fabioberger/humble/view"
	"github.com/fabioberger/recall-frontend/go/models"
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/console"
	"honnef.co/go/js/dom"
)

type App struct {
	humble.Identifier
}

const (
	logOutSelector = "button#sign-out"
)

var (
	doc      = dom.GetWindow().Document()
	elements = struct {
		body dom.Element
	}{}
)

func (a *App) RenderHTML() string {
	return fmt.Sprintf(`
	<nav class="navbar navbar-default">
		  <div class="container-fluid">
		    <div class="navbar-header">
		      <a class="navbar-brand" href="/#/">
		        <img alt="Brand" src="/images/brain.png">
		      </a>
		    </div>
		    <div class="navbar-right collapse navbar-collapse">
		    	<div id="loggedOut">
				    <a href="/#/signup"><button type="button" class="btn btn-default navbar-btn">Sign Up</button></a>
				    <a href="/#/login"><button type="button" class="btn btn-default navbar-btn">Log in</button></a>
				   </div>
				   <div id="loggedIn">
				   		<button type="button" class="btn btn-default navbar-btn" id="sign-out">Sign Out</button>
				   		<a href="/#/profile">
				   			<button type="button" class="btn btn-default navbar-btn">Profile</button>
				   		</a>
				   </div>
		    </div>
		  </div>
		</nav>
    <div id="container" class="container">
    </div>
	<script src="js/app.js"></script>`)
}

func (v *App) OnLoad() error {
	// Check to see if user logged in or not from cookies
	exists := cookieExists("isLoggedIn")
	if exists {
		doc.GetElementByID("loggedOut").SetAttribute("style", "display: none;")
	} else {
		doc.GetElementByID("loggedIn").SetAttribute("style", "display: none;")
	}

	if err := view.AddListener(v, logOutSelector, "click", v.logOutClick); err != nil {
		return err
	}
	return nil
}

func (v *App) logOutClick(event dom.Event) {
	s := &models.Session{}
	if err := model.Delete(s); err != nil {
		console.Log("Hit Error")
		console.Log(err)
	}
	htmlEl := doc.(dom.HTMLDocument)
	htmlEl.SetCookie("isLoggedIn==; expires=Thu, 01 Jan 1970 00:00:01 GMT;")
	htmlEl.SetCookie("recall_session==; expires=Thu, 01 Jan 1970 00:00:01 GMT;")
	setVisitorMenu()
	js.Global.Get("location").Set("hash", "/")
}

func (v *App) OuterTag() string {
	return "div"
}

// Utility function (move to own file?)
// Checks if a particular cookie is set
func cookieExists(name string) bool {
	// var name string
	htmlEl := doc.(dom.HTMLDocument)
	cookieStr := htmlEl.Cookie()
	if cookieStr == "" {
		return false
	}
	cookiePairs := strings.Split(cookieStr, "; ")
	for _, c := range cookiePairs {
		cookie := strings.Split(c, "=")
		if cookie[0] == name {
			return true
		}
	}
	return false
}
