package gen

import (
	"errors"
	"fmt"

	"github.com/StevenZack/gengo/tool"
	"github.com/StevenZack/tools/strToolkit"

	"github.com/StevenZack/tools/fileToolkit"
)

var verbosely bool

func SetVerbosely(b bool) {
	verbosely = b
}

func Gen(args []string) {
	var pkgPath string
	var e error

	if args == nil {
		pkgPath, e = fileToolkit.GetCurrentPkgPath()
		if e != nil {
			fmt.Println(`GetCurrentPkgPath error :`, e)
			return
		}
	} else {
		pkgPath = args[0]
	}

	e = execute(pkgPath)
	if e != nil {
		fmt.Println(`execute error :`, e)
		return
	}
}

func execute(pkgPath string) error {
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
		e := tool.ParseFile(filePath)
	}
	return nil
}

func log(args ...interface{}) {
	if verbosely {
		fmt.Println(args...)
	}
}
