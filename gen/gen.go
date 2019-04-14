package gen

import (
	"fmt"

	"github.com/StevenZack/tools/strToolkit"

	"github.com/StevenZack/tools/fileToolkit"
)

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
	gengoPath := fileToolkit.GetGOPATH() + "src/" + strToolkit.Getrpath(pkgPath)
	if !fileToolkit.IsDirExists(gengoPath) {
		return nil
	}
	fmt.Println("execute:", gengoPath)
	return nil
}
