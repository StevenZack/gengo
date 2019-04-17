package gen

import (
	"encoding/json"
	"os"
	"reflect"
	"runtime"
	"strings"

	"github.com/StevenZack/tools/fileToolkit"
	"github.com/StevenZack/tools/strToolkit"
)

// GengoStruct infers structs in target .go file
type GengoStruct struct {
	PreCompilers []PreCompiler
	GengoTags    []string
	OutputPkgs   []string

	StructPkg string
	FilePath  string
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
	filename := strings.ToLower(g.Name) + "_" + g.PreCompilers[precompilerIndex].PkgName + ".go"
	if len(g.OutputPkgs) > precompilerIndex {
		output := genOutputPath(g.OutputPkgs[precompilerIndex], dir)
		log("\tgetOutputPath(), out:", output)
		return output, nil
	}
	return strToolkit.Getrpath(dir) + filename, nil
}

func genOutputPath(pkg, dir string) string {
	log("\tgetOutputPath() :", pkg, dir)
	sep := string(os.PathSeparator)
	if strings.Contains(pkg, sep) {
		if strings.HasPrefix(pkg, "."+sep) || strings.HasPrefix(pkg, ".."+sep) { // relative path
			return strToolkit.Getrpath(dir) + pkg
		} else if strings.HasPrefix(pkg, sep) || runtime.GOOS == "windows" && strings.Contains(pkg, ":"+sep) { // absolute path
			return pkg
		}
		// go path
		return fileToolkit.GetGOPATH() + "src/" + pkg
	}
	// cur dir
	return strToolkit.Getrpath(dir) + pkg
}
