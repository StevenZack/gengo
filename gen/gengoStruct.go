package gen

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/StevenZack/tools/fileToolkit"
	"github.com/StevenZack/tools/strToolkit"
)

// GengoStruct infers structs in target .go file
type GengoStruct struct {
	PreCompilers []PreCompiler

	StructPkg string
	FilePath  string
	GengoTags  []string
	Name      string
}

// PreCompiler infers a preCompiler
type PreCompiler struct {
	Pkg     string
	PkgName string
}

// Field infers GengoStruct's field
type Field struct {
	Name string
	Kind string
	Tag  reflect.StructTag
}

// GetInfoStr returns GengoStruct's basic infomations , used for logging
func (g *GengoStruct) GetInfoStr() string {
	str := "{}"
	b, e := json.Marshal(g)
	if e == nil {
		str = string(b)
	}
	return g.Name + ":" + str
}

// GetGengoFileOutputPath generate output file path
func (g *GengoStruct) GetGengoFileOutputPath(precompilerIndex int) (string, error) {
	dir, e := fileToolkit.GetDirOfFile(g.FilePath)
	if e != nil {
		return "", e
	}
	return strToolkit.Getrpath(dir) + strings.ToLower(g.Name) + "_" + g.PreCompilers[precompilerIndex].PkgName + ".go", nil
}
