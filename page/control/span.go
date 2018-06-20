package control

import (
	"github.com/spekary/goradd/html"
	"github.com/spekary/goradd/page"
	"github.com/spekary/goradd/page/control/control_base"
)

type SpanI interface {
	control_base.PanelI
}

// Span is a Goradd control that is a basic "span" wrapper. Use it to style and listen to events on a span. It
// can also be used as the basis for more advanced javascript controls.
type Span struct {
	control_base.Panel
}

func NewSpan(parent page.ControlI) *Span {
	p := &Span{}
	p.Init(p, parent)
	return p
}

func (c *Span) Init(self SpanI, parent page.ControlI) {
	c.Panel.Init(self, parent)
	c.Tag = "span"
}

func (c *Span) DrawingAttributes() *html.Attributes {
	a := c.Control.DrawingAttributes()
	a.SetDataAttribute("grctl", "span")
	return a
}