package control

import (
	"context"
	"github.com/goradd/goradd/pkg/bootstrap/config"
	"github.com/goradd/goradd/pkg/html"
	"github.com/goradd/goradd/pkg/page"
	"github.com/goradd/goradd/pkg/page/control"
)

type RadioButtonI interface {
	control.RadioButtonI
}

type RadioButton struct {
	control.RadioButton
	inline bool
}

func NewRadioButton(parent page.ControlI, id string) *RadioButton {
	c := &RadioButton{}
	c.Self = c
	c.Init(parent, id)
	config.LoadBootstrap(c.ParentForm())
	return c
}

func (c *RadioButton) this() RadioButtonI {
	return c.Self.(RadioButtonI)
}

func (c *RadioButton) SetInline(v bool) *RadioButton {
	c.inline = v
	return c
}

func (c *RadioButton) DrawingAttributes(ctx context.Context) html.Attributes {
	a := c.RadioButton.DrawingAttributes(ctx)
	a.SetDataAttribute("grctl", "bs-radio")
	a.AddClass("form-check-input")

	if c.Text() == "" {
		a.AddClass("position-static")
	}
	return a
}

func (c *RadioButton) GetDrawingLabelAttributes() html.Attributes {
	a := c.RadioButton.GetDrawingLabelAttributes()
	a.AddClass("form-check-label")
	return a
}

func (c *RadioButton) DrawTag(ctx context.Context) (ctrl string) {
	h := c.RadioButton.DrawTag(ctx)
	checkWrapperAttributes := html.NewAttributes().
		AddClass("form-check").
		SetDataAttribute("grel", c.ID()) // make sure the entire control gets removed
	if c.inline {
		checkWrapperAttributes.AddClass("form-check-inline")
	}
	return html.RenderTag("div", checkWrapperAttributes, h)
}

type RadioButtonCreator struct {
	// ID is the id of the control
	ID string
	// Text is the text of the label displayed right next to the checkbox.
	Text string
	// Checked will initialize the checkbox in its checked state.
	Checked bool
	// LabelMode specifies how the label is drawn with the checkbox.
	LabelMode html.LabelDrawingMode
	// LabelAttributes are additional attributes placed on the label tag.
	LabelAttributes html.Attributes
	// SaveState will save the value of the checkbox and restore it when the page is reentered.
	SaveState bool
	// Group is the name of the group that the button belongs to
	Group string
	// Inline draws the radio group inline. Specify this when drawing the
	// radio button inline or in an inline FormGroup.
	Inline bool
	page.ControlOptions
}

// Create is called by the framework to create a new control from the Creator. You
// do not normally need to call this.
func (c RadioButtonCreator) Create(ctx context.Context, parent page.ControlI) page.ControlI {
	ctrl := NewRadioButton(parent, c.ID)
	if c.Text != "" {
		ctrl.SetText(c.Text)
	}
	if c.LabelMode != html.LabelDefault {
		ctrl.LabelMode = c.LabelMode
	}
	if c.LabelAttributes != nil {
		ctrl.LabelAttributes().Merge(c.LabelAttributes)
	}
	if c.Group != "" {
		ctrl.SetGroup(c.Group)
	}

	ctrl.ApplyOptions(ctx, c.ControlOptions)
	if c.SaveState {
		ctrl.SaveState(ctx, c.SaveState)
	}
	if c.Inline {
		ctrl.SetInline(c.Inline)
	}
	return ctrl
}

// GetRadioButton is a convenience method to return the radio button with the given id from the page.
func GetRadioButton(c page.ControlI, id string) *RadioButton {
	return c.Page().GetControl(id).(*RadioButton)
}

func init() {
	page.RegisterControl(&RadioButton{})
}