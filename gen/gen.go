package gen

import (
	"errors"
	"fmt"
	"strings"

	"github.com/StevenZack/gengo/tool"
	"github.com/StevenZack/tools/strToolkit"

	"github.com/StevenZack/tools/fileToolkit"
)

var verbosely bool

const genExecutorPkgPath = "github.com/StevenZack/gengo/genexecutor/main.go"

func SetVerbosely(b bool) {
	verbosely = b
}

func Gen(args []string) {
	var pkgPath string
	var e error

	if args == nil {
		pkgPath, e = fileToolkit.GetCurrentPkgPath()
		if e != nil {
			fmt.Errorf("getCurrentPkgPath error :%v", e)
			return
		}
	} else {
		pkgPath = args[0]
	}

	e = compile(pkgPath)
	if e != nil {
		fmt.Errorf("compile %s err:%v", pkgPath, e)
		return
	}
}

func compile(pkgPath string) error {
	absPath := fileToolkit.GetGOPATH() + "src/" + strToolkit.Getunpath(pkgPath)
	if !fileToolkit.IsDirExists(absPath) {
		return errors.New("path:" + pkgPath + " not exists")
	}

	if !fileToolkit.IsDirExists(absPath + "_gengo") {
		return nil
	}

	log(pkgPath)

	list := fileToolkit.GetAllFilesFromFolder(absPath)
	for _, filePath := range list {
		if !strToolkit.EndsWith(filePath, ".go") {
			continue
		}
		if strToolkit.EndsWith(filePath, "_gengo.go") {
			continue
		}
		structs, e := tool.ParseFileGengoStructs(filePath)
		if e != nil {
			return e
		}

		for _, obj := range structs {
			generateExecutor(obj)
		}
	}
	return nil
}

func log(args ...interface{}) {
	if verbosely {
		fmt.Println(args...)
	}
}

func generateExecutor(obj tool.GengoStruct) error {
	path := fileToolkit.GetGOPATH() + "src/" + genExecutorPkgPath
	bakPath := path + ".bak"
	if !fileToolkit.IsFileExists(bakPath) {
		return errors.New("file " + bakPath + " doesn't exists")
	}
	str, e := fileToolkit.ReadFileAll(bakPath)
	if e != nil {
		return e
	}
	str = strings.Replace(str, "github.com/StevenZack/gengo/example/data_gengo", obj.PreCompilerPkg, -1)
	str = strings.Replace(str, "str := data_gengo.Gen(g, t)", "str := "+obj.PreCompilerPkgName+".Gen(g, t)", -1)
	
}
