package views

import (
	"fmt"

	"github.com/fabioberger/humble"
	"github.com/fabioberger/humble/view"
	"github.com/fabioberger/recall-frontend/go/models"
)

type Profile struct {
	humble.Identifier
	Children  []*Reminder
	User      *models.User
	Reminders []*models.Reminder
}

const (
	reminderListSelector = "table#reminder-table tbody"
)

func (v *Profile) RenderHTML() string {
	return fmt.Sprintf(`
		<h1>Welcome %s!</h1>
		<p>Notification Email: %s</p>

		<div class="new-reminder">
			<a href="/#/">
				<button class="btn btn-success">New Reminder</button>
			</a>
		</div>

		<div class="reminders">
			<table class="table table-hover table-bordered" id="reminder-table">
				<tr>
							<th>Reminders</th>
							<th>Created on</th>
							<th>Remove</th>
				</tr>
			</table>
		</div>`,
		v.User.Name, v.User.Email)
}

func (v *Profile) InitChildren(reminders []*models.Reminder) {
	//Create individual todo views
	v.Children = []*Reminder{}
	for _, r := range reminders {
		reminderView := &Reminder{
			Reminder: r,
			Parent:   v,
		}
		v.addChild(reminderView)
	}
}

func (v *Profile) addChild(reminderView *Reminder) {
	v.Children = append(v.Children, reminderView)
}

func (v *Profile) removeChild(reminderView *Reminder) {
	for i, child := range v.Children {
		if child.Id == reminderView.Id {
			v.Children = append(v.Children[:i], v.Children[i+1:]...)
		}
	}
}

func (v *Profile) OnLoad() error {
	// Add each child view to the DOM
	for _, childView := range v.Children {
		view.AppendToParentHTML(childView, reminderListSelector)
	}
	return nil
}

func (v *Profile) OuterTag() string {
	return "div"
}
