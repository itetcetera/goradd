//** This file was code generated by got. ***

package template

import (
	"bytes"
	"fmt"
	"goradd/config"

	"github.com/spekary/goradd/codegen/generator"
	"github.com/spekary/goradd/orm/db"
)

func init() {
	t := ConnectorBaseTemplate{
		generator.Template{
			Overwrite: true,
			TargetDir: config.LocalDir + "/gen",
		},
	}
	generator.AddTableTemplate(&t)
}

type ConnectorBaseTemplate struct {
	generator.Template
}

func (n *ConnectorBaseTemplate) FileName(key string, t *db.TableDescription) string {
	return n.TargetDir + "/" + key + "/connector/" + t.GoName + ".go"
}

func (n *ConnectorBaseTemplate) GenerateTable(codegen generator.Codegen, dd *db.DatabaseDescription, t *db.TableDescription, buf *bytes.Buffer) {
	//var privateName = util.LcFirst(t.GoName)
	//connector.tmpl

	// The master template for the connector classes

	buf.WriteString(`package connector

// This file is code generated. Do not edit.

import (
	"goradd/gen/`)

	buf.WriteString(fmt.Sprintf("%v", t.DbKey))

	buf.WriteString(`/model/node"
	"github.com/spekary/goradd/orm/db"
	"github.com/spekary/goradd/orm/query"
	"context"
	"fmt"
	. "github.com/spekary/goradd/orm/op"
	//"./node"
	"github.com/spekary/goradd/datetime"
	"github.com/spekary/goradd/util"
)

`)

}

func (n *ConnectorBaseTemplate) Overwrite() bool {
	return n.Template.Overwrite
}
