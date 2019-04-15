package gen

import (
	"encoding/json"
	"reflect"

	"github.com/StevenZack/tools/fileToolkit"
	"github.com/StevenZack/tools/strToolkit"
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

func (g *GengoStruct) GetGengoFileOutputPath() (string, error) {
	name, e := fileToolkit.GetNameOfPath(g.FilePath)
	if e != nil {
		return "", e
	}
	dir, e := fileToolkit.GetDirOfFile(g.FilePath)
	if e != nil {
		return "", e
	}
	nameWithoutGo := name[:len(name)-len(".go")]
	return strToolkit.Getrpath(dir) + nameWithoutGo + "_gengo.go", nil
}
