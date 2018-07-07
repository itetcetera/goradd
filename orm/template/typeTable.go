//** This file was code generated by got. ***

package template

import (
	"bytes"
	"goradd/config"
	"strconv"
	"strings"

	"github.com/gedex/inflector"
	"github.com/knq/snaker"
	"github.com/spekary/goradd/codegen/generator"
	"github.com/spekary/goradd/orm/db"
	"github.com/spekary/goradd/orm/query"
)

func init() {
	t := TypeTableTemplate{
		generator.Template{
			Overwrite: true,
			TargetDir: config.LocalDir + "/gen",
		},
	}
	generator.AddTypeTableTemplate(&t)
}

type TypeTableTemplate struct {
	generator.Template
}

func (n *TypeTableTemplate) FileName(key string, tt *db.TypeTableDescription) string {
	return n.TargetDir + "/" + key + "/model/" + tt.GoName + ".base.go"
}

func (n *TypeTableTemplate) GenerateTypeTable(codegen generator.Codegen, dd *db.DatabaseDescription, tt *db.TypeTableDescription, buf *bytes.Buffer) {
	//var privateName = util.LcFirst(tt.GoName)

	propName := tt.GoName
	propLetter := strings.ToLower(propName[0:1])
	keyField := tt.FieldNames[0]

	buf.WriteString(`package model

// This file is code generated. Do not edit.

import (
	//"log"
	//"github.com/spekary/goradd/orm/query"
)

const (
`)

	for _, value := range tt.Values {
		key, ok := value[keyField].(uint)
		if !ok {
			key = uint(value[keyField].(int))
		}
		con := tt.Constants[key]

		buf.WriteString(`	`)

		buf.WriteString(propName)

		buf.WriteString(con)

		buf.WriteString(` `)

		buf.WriteString(propName)

		buf.WriteString(` = `)

		buf.WriteString(strconv.FormatUint(uint64(key), 10))

		buf.WriteString(`
`)

	}

	buf.WriteString(`)

type `)

	buf.WriteString(propName)

	buf.WriteString(` uint

func (`)

	buf.WriteString(propLetter)

	buf.WriteString(` `)

	buf.WriteString(propName)

	buf.WriteString(`) String() string {
	switch `)

	buf.WriteString(propLetter)

	buf.WriteString(` {
	case 0: return ""
`)
	for _, value := range tt.Values {
		buf.WriteString(`	case `)

		buf.WriteString(generator.AsConstant(value[keyField], query.ColTypeUnsigned))

		buf.WriteString(`: return `)

		buf.WriteString(generator.AsConstant(value[tt.FieldNames[1]], query.ColTypeString))

		buf.WriteString(`
`)
	}
	buf.WriteString(`	default: panic("Index out of range")
	}
	return "" // prevent warning
}

`)

	for i, fieldName := range tt.FieldNames[2:] {
		typ := tt.FieldTypes[fieldName]
		typeName := string(tt.FieldTypes[fieldName])
		title := snaker.SnakeToCamel(fieldName)

		buf.WriteString(`func (`)

		buf.WriteString(propLetter)

		buf.WriteString(` `)

		buf.WriteString(propName)

		buf.WriteString(`) `)

		buf.WriteString(title)

		buf.WriteString(`() `)

		buf.WriteString(typeName)

		buf.WriteString(` {
	switch `)

		buf.WriteString(propLetter)

		buf.WriteString(` {
	case 0: return `)

		buf.WriteString(typ.DefaultValue())

		buf.WriteString(`
`)
		for _, value := range tt.Values {
			buf.WriteString(`	case `)

			buf.WriteString(generator.AsConstant(value[keyField], query.ColTypeUnsigned))

			buf.WriteString(`: return `)

			buf.WriteString(generator.AsConstant(value[tt.FieldNames[i+2]], typ))

			buf.WriteString(`
`)
		}
		buf.WriteString(`	default: panic("Index out of range")
	}
	return `)

		buf.WriteString(typ.DefaultValue())

		buf.WriteString(` // prevent warning
}
`)

	}

	for _, fieldName := range tt.FieldNames[1:] {
		typeName := string(tt.FieldTypes[fieldName])
		varName := inflector.Pluralize(fieldName)
		title := inflector.Pluralize(snaker.SnakeToCamel(fieldName))

		buf.WriteString(`func `)

		buf.WriteString(propName)

		buf.WriteString(title)

		buf.WriteString(`() []`)

		buf.WriteString(typeName)

		buf.WriteString(` {
	`)

		buf.WriteString(varName)

		buf.WriteString(` := make([]`)

		buf.WriteString(typeName)

		buf.WriteString(`, `)

		buf.WriteString(strconv.Itoa(len(tt.Values) + 1))

		buf.WriteString(`)
	// 0 item will be a blank
`)
		for i, value := range tt.Values {
			buf.WriteString(`	`)

			buf.WriteString(varName)

			buf.WriteString(`[`)

			buf.WriteString(strconv.Itoa(i + 1))

			buf.WriteString(`] = `)

			buf.WriteString(generator.AsConstant(value[fieldName], tt.FieldTypes[fieldName]))

			buf.WriteString(`
`)
		}
		buf.WriteString(`	return `)

		buf.WriteString(varName)

		buf.WriteString(`
}

`)

	}

}

func (n *TypeTableTemplate) Overwrite() bool {
	return n.Template.Overwrite
}
