package tool

import (
	"github.com/StevenZack/tools/fileToolkit"
)

func IsGoPkg(pkgPath string) bool {
	if pkgPath == "" {
		return false
	}
	gopath := fileToolkit.GetGOPATH()
	return fileToolkit.IsDirExists(gopath + "src/" + pkgPath)
}
