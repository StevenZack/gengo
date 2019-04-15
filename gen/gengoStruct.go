package gen

import (
	"encoding/json"
	"reflect"
)

type GengoStruct struct {
	PreCompilerPkg     string
	PreCompilerPkgName string

	StructPkg string
	FilePath  string
	GengoTag  string
	Name      string
}
type Field struct {
	Name string
	Kind string
	Tag  reflect.StructTag
}

func (g *GengoStruct) GetInfoStr() string {
	str := "{}"
	b, e := json.Marshal(g)
	if e == nil {
		str = string(b)
	}
	return g.Name + ":" + str
}
