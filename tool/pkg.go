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

func GetPkgNameFromPkg(gopkg string) (string, error) {
	dir := fileToolkit.GetGOPATH() + "src/" + gopkg
	gofile, e := GetFirstGoFile(dir)
	if e != nil {
		return "", e
	}
	firstLine, e := ReadFirstLine(gofile)
	if e != nil {
		return "", e
	}
	pkg, e := ReadPkgFromLine(firstLine)
	if e != nil {
		return "", e
	}
	return pkg, nil
}

func GetPkgFromFilePath(filePath string) (string, error) {
	dir, e := fileToolkit.GetDirOfFile(filePath)
	if e != nil {
		return "", e
	}
	fileToolkit.GetGOPATH()
}
