package gen

import (
	"errors"
	"fmt"
	"os"

	"github.com/StevenZack/tools/fileToolkit"
	"github.com/StevenZack/tools/ioToolkit"
	"github.com/StevenZack/tools/strToolkit"
)

func compile(pkgPath string) error {
	if !fileToolkit.IsGoPathPkg(pkgPath) {
		return nil
	}

	absPath := fileToolkit.GetGOPATH() + "src/" + strToolkit.Getunpath(pkgPath)
	list, e := fileToolkit.RangeFilesInDir(absPath)
	if e != nil {
		return e
	}
	log("compiling ", pkgPath)
	for _, filePath := range list {
		if !strToolkit.EndsWith(filePath, ".go") {
			continue
		}
		if strToolkit.EndsWith(filePath, "_gengo.go") {
			continue
		}

		allImports, e := fileToolkit.GetAllImports(filePath)
		if e != nil {
			return errors.New(filePath + " getAllImports failed:" + e.Error())
		}
		for _, imp := range allImports {
			if !fileToolkit.IsGoPathPkg(imp) {
				continue
			}
			e := compile(imp)
			if e != nil {
				fmt.Println("compile "+imp+" error :", e)
			}
		}

		log("\tparsing file:", filePath)
		structs, e := ParseFileGengoStructs(filePath)
		if e != nil {
			return errors.New(filePath + " parseFileGengoStructs failed:" + e.Error())
		}

		for _, obj := range structs {
			outputFile, e := obj.GetGengoFileOutputPath()
			log("\t\t", obj.GetInfoStr())
			if e != nil {
				return e
			}

			e = os.Remove(outputFile)
			if e != nil {
				log("\t\tos.Remove file failed:", outputFile)
			}

			e = generateExecutor(obj)
			if e != nil {
				return errors.New(filePath + " gen executor failed:" + e.Error())
			}

			// run : genexecutor
			e = ioToolkit.RunAttachedCmd("genexecutor", outputFile, obj.GengoTag)
			if e != nil {
				return errors.New(filePath + " " + e.Error() + " . Did you forget to add GOPATH/bin to $PATH environment variable ?")
			}
		}
	}
	return nil
}
