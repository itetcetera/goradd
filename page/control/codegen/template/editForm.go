//** This file was code generated by got. ***

package template

import (
	"bytes"
	"fmt"
	"goradd/config"
	"strings"

	"github.com/spekary/goradd/codegen/generator"
	"github.com/spekary/goradd/orm/db"
)

func init() {
	t := EditFormTemplate{
		generator.Template{
			Overwrite: true,
			TargetDir: config.LocalDir + "/form",
		},
	}
	generator.AddTableTemplate(&t)
}

type EditFormTemplate struct {
	generator.Template
}

func (n *EditFormTemplate) FileName(key string, t generator.TableType) string {
	return n.TargetDir + "/" + strings.Title(key) + t.GoName + "EditForm.go"
}

func (n *EditFormTemplate) GenerateTable(codegen generator.Codegen, dd *db.DatabaseDescription, t generator.TableType, buf *bytes.Buffer) {
	//editForm.tmpl

	// The master template for the EditForm classes

	var key string = strings.Title(dd.DbKey)
	var formName = key + t.GoName + "EditForm"

	buf.WriteString(`package form

import (
    "github.com/spekary/goradd/page"
     .  "github.com/spekary/goradd/page/control"
     "goradd/gen/`)

	buf.WriteString(fmt.Sprintf("%v", dd.DbKey))

	buf.WriteString(`/panel"
    "context"
    "github.com/spekary/goradd/page/action"
	"fmt"
)

const `)

	buf.WriteString(key)

	buf.WriteString(t.GoName)

	buf.WriteString(`EditPath = "/form/`)

	buf.WriteString(key)

	buf.WriteString(`/`)

	buf.WriteString(t.GoName)

	buf.WriteString(`Edit"
const `)

	buf.WriteString(key)

	buf.WriteString(t.GoName)

	buf.WriteString(`EditID = "`)

	buf.WriteString(formName)

	buf.WriteString(`"
const `)

	buf.WriteString(t.GoName)

	buf.WriteString(`Singular = "`)

	buf.WriteString(t.GoName)

	buf.WriteString(`"
const `)

	buf.WriteString(t.GoName)

	buf.WriteString(`Plural = "`)

	buf.WriteString(t.GoPlural)

	buf.WriteString(`"

const (
    `)

	buf.WriteString(t.GoName)

	buf.WriteString(`SaveAction = iota + 1
    `)

	buf.WriteString(t.GoName)

	buf.WriteString(`CancelAction
    `)

	buf.WriteString(t.GoName)

	buf.WriteString(`DeleteAction
)


// The `)

	buf.WriteString(formName)

	buf.WriteString(` is a form wrapper for the corresponding edit panel.
// To edit it, make a copy and move it out of this package and into another
type `)

	buf.WriteString(formName)

	buf.WriteString(` struct {
    FormBase
    EditPanel *panel.`)

	buf.WriteString(t.GoName)

	buf.WriteString(`Edit
    SaveButton *Button
    CancelButton *Button
    DeleteButton *Button
}

func New`)

	buf.WriteString(formName)

	buf.WriteString(`(ctx context.Context) page.FormI {
    f := new(`)

	buf.WriteString(formName)

	buf.WriteString(`)
    f.Init(ctx, f, `)

	buf.WriteString(key)

	buf.WriteString(t.GoName)

	buf.WriteString(`EditPath, `)

	buf.WriteString(key)

	buf.WriteString(t.GoName)

	buf.WriteString(`EditID)
    return f
}

func (f *`)

	buf.WriteString(formName)

	buf.WriteString(`)Init(ctx context.Context, self page.FormI, path string, id string) {
    f.FormBase.Init(ctx, self, path, id)

	f.AddRelatedFiles()
	f.CreateControls(ctx)
	f.LoadControls(ctx)
}

func (f *`)

	buf.WriteString(formName)

	buf.WriteString(`) AddReleatedFiles() {
    f.FormBase.AddRelatedFiles()

    // Add additional javascript, css and other files here
}

func (f *`)

	buf.WriteString(formName)

	buf.WriteString(`) CreateControls(ctx context.Context) {
	f.EditPanel = panel.New`)

	buf.WriteString(t.GoName)

	buf.WriteString(`EditPanel(f, "edit-panel")
	f.SaveButton = NewButton(f, "save-button")
	f.SaveButton.OnClick(action.Ajax(f.ID(), `)

	buf.WriteString(t.GoName)

	buf.WriteString(`SaveAction))
	f.CancelButton = NewButton(f, "cancel-button")
	f.CancelButton.OnClick(action.Ajax(f.ID(), `)

	buf.WriteString(t.GoName)

	buf.WriteString(`CancelAction))
	f.DeleteButton = NewButton(f, "delete-button")
	f.DeleteButton.OnClick(
	    action.Confirm(fmt.Sprintf(f.T("Are you sure you want to delete this %s"), `)

	buf.WriteString(t.GoName)

	buf.WriteString(`Singular)),
	    action.Ajax(f.ID(), `)

	buf.WriteString(t.GoName)

	buf.WriteString(`DeleteAction),
	    )
}

func (f *`)

	buf.WriteString(formName)

	buf.WriteString(`) LoadControls(ctx context.Context) {
    if id, ok := page.GetContext(ctx).FormValue("id"); ok {
    	f.EditPanel.Load(ctx, id)
        f.DeleteButton.SetVisible(true)
        f.Page().SetTitle(fmt.Sprintf(f.T("Edit %s"), `)

	buf.WriteString(t.GoName)

	buf.WriteString(`Singular))
    } else {
        f.EditPanel.Load(ctx, "")
        f.DeleteButton.SetVisible(false)
        f.Page().SetTitle(fmt.Sprintf(f.T("New %s"), `)

	buf.WriteString(t.GoName)

	buf.WriteString(`Singular))
    }
}

func (f *`)

	buf.WriteString(formName)

	buf.WriteString(`) Action(ctx context.Context, a page.ActionParams) {
	switch a.ID {
	case `)

	buf.WriteString(t.GoName)

	buf.WriteString(`SaveAction:
	    f.EditPanel.Save(ctx)
	    f.ReturnToPrevious()
    case `)

	buf.WriteString(t.GoName)

	buf.WriteString(`CancelAction:
        f.ReturnToPrevious()
    case `)

	buf.WriteString(t.GoName)

	buf.WriteString(`DeleteAction:
        f.EditPanel.Delete(ctx)
        f.ReturnToPrevious()
    }
}

func (f *`)

	buf.WriteString(formName)

	buf.WriteString(`) ReturnToPrevious() {
	f.ChangeLocation("/form/`)

	buf.WriteString(t.GoName)

	buf.WriteString(`List")
}

func init() {
	page.RegisterPage(`)

	buf.WriteString(key)

	buf.WriteString(t.GoName)

	buf.WriteString(`EditPath,  New`)

	buf.WriteString(formName)

	buf.WriteString(`, `)

	buf.WriteString(key)

	buf.WriteString(t.GoName)

	buf.WriteString(`EditID)
}


`)

}

func (n *EditFormTemplate) Overwrite() bool {
	return n.Template.Overwrite
}
