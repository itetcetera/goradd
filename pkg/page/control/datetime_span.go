package control

import (
	"github.com/goradd/goradd/pkg/config"
	"github.com/goradd/goradd/pkg/page"
	"bytes"
	"context"
	"github.com/goradd/goradd/pkg/datetime"
	"time"
)

// DateTimeSpan is a span that displays a datetime value as static text. This is a typical default control to use
// for a timestamp in the database.
type DateTimeSpan struct {
	Span
	format string
	value datetime.DateTime
}

// NewDateTimeSpan create a new DateTimeSpan.
func NewDateTimeSpan(parent page.ControlI, id string) *DateTimeSpan {
	s := new(DateTimeSpan)
	s.Init(s, parent, id)
	return s
}

// Init is called by subclasses to initialize the parent. You do not normally need to call this.
func (s *DateTimeSpan) Init(self page.ControlI, parent page.ControlI, id string) {
	s.Span.Init(self, parent, id)
	s.format = config.DefaultDateTimeFormat
}

// SetValue sets the value display. You can set the value to a datetime.DateTime, a time.Time,
// or a string that can be parsed by the format string.
func (s *DateTimeSpan) SetValue(v interface{}) {
	switch v2 := v.(type) {
	case datetime.DateTime:
		s.SetDateTime(v2)
	case time.Time:
		s.SetDateTime(datetime.NewDateTime(v2))
	case string:
		d,err := datetime.Parse(s.format, v2)
		if err != nil {
			panic(err)
		}
		s.SetDateTime(d)
	}
}

// SetDateTime sets the value to a datetime.DateTime.
func (s *DateTimeSpan) SetDateTime(d datetime.DateTime){
	s.value = d
	s.Refresh()
}

// Value returns the value as a datetime.DateTime object.
func (s *DateTimeSpan) Value() datetime.DateTime {
	return s.value
}

// SetFormat sets the format string. This should be a time.TimeFormat string described at
// https://golang.org/pkg/time/#Time.Format
func (s *DateTimeSpan) SetFormat(format string)  *DateTimeSpan {
	s.format = format
	s.Refresh()
	return s
}

// DrawInnerHtml is called by the framework to draw the inner html of the span.
func (s *DateTimeSpan) DrawInnerHtml(ctx context.Context, buf *bytes.Buffer) error {
	buf.WriteString(s.value.Format(s.format))
	// TODO: Internationalize this. Golang does not currently have date/time localization support,
	// as in translation of month and weekday strings, but it does support arbitrary timezones.
	// The other option is to let JavaScript format it, though that limits you to formatting in
	// local time or UTC. JavaScript does not have a means to specify the timezone that is well supported.
	// However, JavaScript will translate month and weekday names to the local language.
	return nil
}
