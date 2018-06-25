//** This file was code generated by got. ***

package control

import (
	"bytes"
	"context"
	"fmt"
)

func (t *Tabs) DrawTemplate(ctx context.Context, buf *bytes.Buffer) (err error) {
	var children = t.Children()

	if (children != nil && len(children) > 0) &&
		(t.selectedID == "" || t.Child(t.selectedID) == nil) {
		t.selectedID = children[0].ID() // select first item if nothing is selected
	}

	buf.WriteString(`
<ul class="nav nav-tabs" id="myTab" role="tablist">
`)
	for _, child := range children {
		buf.WriteString(`  <li class="nav-item">
    <a class="nav-link `)
		if child.ID() == t.selectedID {
			buf.WriteString(`active`)
		}

		buf.WriteString(`" id="`)

		buf.WriteString(fmt.Sprintf("%v", child.ID()))

		buf.WriteString(`_tab" data-toggle="tab" href="#`)

		buf.WriteString(fmt.Sprintf("%v", child.ID()))

		buf.WriteString(`" role="tab" aria-controls="`)

		buf.WriteString(fmt.Sprintf("%v", child.ID()))

		buf.WriteString(`" aria-selected="true">`)

		buf.WriteString(fmt.Sprintf("%v", child.Label()))

		buf.WriteString(`</a>
  </li>
`)
	}

	buf.WriteString(`</ul>
<div class="tab-content" id="myTabContent">
`)
	for _, child := range children {
		child.AddClass("tab-pane")
		if child.ID() == t.selectedID {
			child.AddClass("active")
		}
		child.SetAttribute("role", "tabpanel")
		child.SetAttribute("aria-labelledby", child.ID()+"_tab")

		buf.WriteString(`
`)

		{
			err := child.Draw(ctx, buf)
			if err != nil {
				return err
			}
		}

		buf.WriteString(`
`)
	}

	buf.WriteString(`</div>

`)

	return
}
