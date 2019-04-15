package gen

import (
	"fmt"

	"github.com/StevenZack/tools/strToolkit"

	"github.com/StevenZack/tools/fileToolkit"
)

var verbosely bool

const genExecutorPkgPath = "github.com/StevenZack/gengo/genexecutor"

// SetVerbosely set verbosely
func SetVerbosely(b bool) {
	verbosely = b
}

// Gen implements the gengo-gen action
func Gen(args []string) {
	var pkgPath string
	var e error

	if args == nil {
		pkgPath, e = fileToolkit.GetCurrentPkgPath()
		if e != nil {
			fmt.Printf("getCurrentPkgPath error :%v", e)
			return
		}
	} else {
		pkgPath = args[0]
	}
	log("target pkgPath =", pkgPath, "\n")

	absPath := fileToolkit.GetGOPATH() + "src/" + strToolkit.Getunpath(pkgPath)
	if !fileToolkit.IsDirExists(absPath) {
		fmt.Println("path:" + pkgPath + " not exists")
		return
	}

	e = compile(pkgPath)
	if e != nil {
		fmt.Printf("compile %s err:%v", pkgPath, e)
		return
	}
}

// log print logs if variable verbosely is true
func log(args ...interface{}) {
	if verbosely {
		fmt.Println(args...)
	}
}
