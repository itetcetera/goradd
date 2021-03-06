package control

import (
	"bytes"
	"context"
	"encoding/gob"
	"github.com/goradd/goradd/pkg/html"
	"github.com/stretchr/testify/assert"
	"testing"
)


type pagedTableTestForm struct {
	FormBase
}

func (*pagedTableTestForm) 	TableRowAttributes(row int, data interface{}) html.Attributes {
	return html.NewAttributes().AddAttributeValue("a", "b")
}

func (*pagedTableTestForm) 	TableHeaderRowAttributes(row int) html.Attributes {
	return html.NewAttributes().AddAttributeValue("c", "d")
}

func (*pagedTableTestForm) 	TableFooterRowAttributes(row int) html.Attributes {
	return html.NewAttributes().AddAttributeValue("e", "f")
}

func (*pagedTableTestForm) 	BindData(ctx context.Context, s DataManagerI) {

}

func TestPagedTable_Serialize(t *testing.T) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	f := &pagedTableTestForm{}
	f.Self = f
	f.FormBase.Init(context.Background(), "MockFormId")

	f.AddControls(context.Background(),
		PagedTableCreator{
			ID:                "table",
			HasColTags:        true,
			Caption:           "This is a table",
			HideIfEmpty:       true,
			HeaderRowCount:    2,
			FooterRowCount:    3,
			RowStylerID:       f.ID(),
			HeaderRowStyler:   f,
			FooterRowStyler:   f,
			DataProvider:      f,
			Sortable:          true,
			SortHistoryLimit:  3,
			OnCellClick:       nil,
			PageSize:          7,
			SaveState:         false, // must have a session to test
			Columns: nil,		// testing columns here will cause circular import
		},
		DataPagerCreator{
			ID:               "dp",
			PagedControl:     "table",
		},
	)

	c := GetPagedTable(f, "table")

	c.Serialize(enc)

	c2 := PagedTable{}
	dec := gob.NewDecoder(&buf)
	c2.Deserialize(dec)

	assert.True(t, c2.Table.renderColumnTags)
	assert.Equal(t, "This is a table", c2.caption)
	assert.Equal(t, 3, c2.sortHistoryLimit)
}
