package generator

import (
	"goradd/config"
	"github.com/spekary/goradd/codegen/connector"
	"github.com/spekary/goradd/orm/db"
	"fmt"
	"goradd/config/codegen"
	"github.com/spekary/goradd/util/types"
	"github.com/spekary/goradd/page"
	"github.com/spekary/goradd/orm/query"
)

func init() {
	if config.Mode == config.AppModeDevelopment {
		connector.RegisterGenerator(Checkbox{})
	}
}

// This structure describes the Checkbox to the connector dialog and code generator
type Checkbox struct {

}

func (d Checkbox) Type() string {
	return "Checkbox"
}

func (d Checkbox) NewFunc() string {
	return "NewCheckbox"
}

func (d Checkbox) Import() string {
	// TODO: Add fmt to the import list
	return "github.com/spekary/goradd/page/control"
}

func (d Checkbox) SupportsColumn(col *db.ColumnDescription) bool {
	if col.GoType == query.ColTypeBool {
		return true
	}
	return false
}

func (d Checkbox) GenerateCreate(namespace string, col *db.ColumnDescription) (s string) {
	s = fmt.Sprintf(
		`	ctrl = %s.NewCheckbox(c.ParentControl, id)
	ctrl.SetLabel("%s")
`, namespace, col.GoName)

	if codegen.DefaultWrapper != "" {
		s += fmt.Sprintf(`	ctrl.With(page.NewWrapper("%s"))
`, codegen.DefaultWrapper)
	}

	return
}

func (d Checkbox) GenerateGet(ctrlName string, objName string, col *db.ColumnDescription) (s string) {
	s = fmt.Sprintf(`c.%s.SetChecked(c.%s.%s())`, ctrlName, objName, col.GoName)
	return
}

func (d Checkbox) GeneratePut(ctrlName string, objName string, col *db.ColumnDescription) (s string) {
	s = fmt.Sprintf(`c.%s.Set%s(c.%s.Checked())`, objName,  col.GoName, ctrlName)
	return
}


func (d Checkbox) ConnectorParams() *types.OrderedMap {
	paramControls := page.ControlConnectorParams()

	return paramControls
}

