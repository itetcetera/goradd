package control

import (
	"encoding/gob"
	"github.com/spekary/goradd/pkg/html"
	"github.com/spekary/goradd/pkg/page"
	"github.com/spekary/goradd/pkg/page/control"
)

type IntegerTextbox struct {
	control.IntegerTextbox
}

func NewIntegerTextbox(parent page.ControlI, id string) *IntegerTextbox {
	t := new (IntegerTextbox)
	t.Init(t, parent, id)
	return t
}


func (t *IntegerTextbox) DrawingAttributes() *html.Attributes {
	a := t.IntegerTextbox.DrawingAttributes()
	a.AddClass("form-control")
	return a
}


func init () {
	gob.RegisterName("bootstrap.integertextbox", new(IntegerTextbox))
}

