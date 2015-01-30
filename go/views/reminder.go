package views

import (
	"fmt"

	"honnef.co/go/js/dom"

	"github.com/fabioberger/humble"
	"github.com/fabioberger/humble/model"
	"github.com/fabioberger/humble/view"
	"github.com/fabioberger/recall-frontend/go/models"
)

type Reminder struct {
	humble.Identifier
	Reminder *models.Reminder
	Parent   *Profile
}

func (r *Reminder) RenderHTML() string {
	return fmt.Sprintf(`
	<tr>
		<td>%s</td>
		<td>%s</td>
		<td class="remove"><a href="javascript:void(0)" class="remove">X</a></td>
	</tr>`,
		r.Reminder.Reminder, r.Reminder.GetDate())
}

func (r *Reminder) OnLoad() error {
	if err := view.AddListener(r, "a.remove", "click", r.deleteClicked); err != nil {
		return err
	}
	return nil
}

func (r *Reminder) deleteClicked(event dom.Event) {
	r.remove()
}

func (r *Reminder) remove() {
	if err := view.Remove(r); err != nil {
		panic(err)
	}
	if err := model.Delete(r.Reminder); err != nil {
		panic(err)
	}
	r.Parent.removeChild(r)
}

func (r *Reminder) OuterTag() string {
	return "tr"
}
