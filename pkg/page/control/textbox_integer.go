package control

import (
	"context"
	"fmt"
	"github.com/goradd/goradd/pkg/page"
	"strconv"
)

type IntegerTextbox struct {
	Textbox
	minValue *int
	maxValue *int
}

func NewIntegerTextbox(parent page.ControlI, id string) *IntegerTextbox {
	t := &IntegerTextbox{}
	t.Init(t, parent, id)
	return t
}

func (i *IntegerTextbox) Init(self TextboxI, parent page.ControlI, id string) {
	i.Textbox.Init(self, parent, id)
	i.ValidateWith(IntValidator{})
}

// SetMinValue creates a validator that makes sure the value of the text box is at least the
// given value. Specify your own error message, or leave the error message blank and a standard error message will
// be presented if the value is not valid.
func (i *IntegerTextbox) SetMinValue(minValue int, invalidMessage string) {
	i.ValidateWith(MinIntValidator{minValue, invalidMessage})
	i.minValue = new(int)
	*i.minValue = minValue
}

// SetMaxValue creates a validator that makes sure the value of the text box is at most the
// given value. Specify your own error message, or leave the error message blank and a standard error message will
// be presented if the value is not valid.
func (i *IntegerTextbox) SetMaxValue(maxValue int, invalidMessage string) {
	i.ValidateWith(MaxIntValidator{maxValue, invalidMessage})
	i.maxValue = new(int)
	*i.maxValue = maxValue
}

func (i *IntegerTextbox) SetValue(v interface{}) *IntegerTextbox {
	i.Textbox.SetValue(v)
	newValue := i.Int()
	if i.minValue != nil && *i.minValue > newValue {
		panic("Setting IntegerTextbox to a value less than minimum value.")
	}
	if i.maxValue != nil && *i.maxValue < newValue {
		panic("Setting IntegerTextbox to a value greater than the maximum value.")
	}
	return i
}

func (i *IntegerTextbox) SetInt(v int) *IntegerTextbox {
	i.Textbox.SetValue(v)
	if i.minValue != nil && *i.minValue > v {
		panic("Setting IntegerTextbox to a value less than minimum value.")
	}
	if i.maxValue != nil && *i.maxValue < v {
		panic("Setting IntegerTextbox to a value greater than the maximum value.")
	}
	return i
}

func (i *IntegerTextbox) Value() interface{} {
	return i.Int()
}

func (i *IntegerTextbox) Int() int {
	t := i.Textbox.Text()
	v, _ := strconv.Atoi(t)
	return v
}

func (i *IntegerTextbox) Int64() int64 {
	t := i.Textbox.Text()
	i64, _ := strconv.ParseInt(t, 10, 0)
	return i64
}

type IntValidator struct {
	Message string
}

func (v IntValidator) Validate(c page.ControlI, s string) (msg string) {
	if s == "" {
		return "" // empty textbox is checked elsewhere
	}
	if _, err := strconv.Atoi(s); err != nil {
		if v.Message == "" {
			return c.T("Please enter an integer.")
		} else {
			return v.Message
		}
	}
	return
}

type MinIntValidator struct {
	MinValue int
	Message  string
}

func (v MinIntValidator) Validate(c page.ControlI, s string) (msg string) {
	if s == "" {
		return "" // empty textbox is checked elsewhere
	}
	if val, _ := strconv.Atoi(s); val < v.MinValue {
		if v.Message == "" {
			return fmt.Sprintf(c.ΩT("Enter at least %d"), v.MinValue)
		} else {
			return v.Message
		}
	}
	return
}

type MaxIntValidator struct {
	MaxValue int
	Message  string
}

func (v MaxIntValidator) Validate(c page.ControlI, s string) (msg string) {
	if s == "" {
		return "" // empty textbox is checked elsewhere
	}
	if val, _ := strconv.Atoi(s); val < v.MaxValue {
		if v.Message == "" {
			return fmt.Sprintf(c.ΩT("Enter at most %d"), v.MaxValue)
		} else {
			return v.Message
		}
	}
	return
}

type IntegerLimit struct {
	Value int
	InvalidMessage string
}


type IntegerTextboxCreator struct {
	ID string
	Placeholder string
	Type string
	MinLength int
	MaxLength int
	ColumnCount int
	RowCount int
	ReadOnly bool
	SaveState bool
	MinValue *IntegerLimit
	MaxValue *IntegerLimit
	// Value is the initial value of the textbox. Often its best to load the value in a separate Load step after creating the control.
	Value interface{}

	page.ControlOptions
}

func (t IntegerTextboxCreator) Create(ctx context.Context, parent page.ControlI) page.ControlI {
	ctrl := NewIntegerTextbox(parent, t.ID)
	if t.Placeholder != "" {
		ctrl.SetPlaceholder(t.Placeholder)
	}
	if t.Type != "" {
		ctrl.typ = t.Type
	}
	ctrl.minLength = t.MinLength
	ctrl.maxLength = t.MaxLength
	ctrl.rowCount = t.RowCount
	ctrl.columnCount = t.ColumnCount
	ctrl.readonly = t.ReadOnly
	if t.MinValue != nil {
		ctrl.SetMinValue(t.MinValue.Value, t.MinValue.InvalidMessage)
	}
	if t.MaxValue != nil {
		ctrl.SetMinValue(t.MaxValue.Value, t.MaxValue.InvalidMessage)
	}
	if t.Value != nil {
		ctrl.SetValue(t.Value)
	}

	ctrl.ApplyOptions(t.ControlOptions)
	if t.SaveState {
		ctrl.SaveState(ctx, t.SaveState)
	}
	return ctrl
}
