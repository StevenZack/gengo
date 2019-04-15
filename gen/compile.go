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
			log("\t\t", obj.GetInfoStr())
			for preIndex := range obj.PreCompilers {
				outputFile, e := obj.GetGengoFileOutputPath(preIndex)
				if e != nil {
					return e
				}

				e = os.Remove(outputFile)
				if e != nil {
					log("\t\tos.Remove file failed:", outputFile)
				}

				e = generateExecutor(obj, preIndex)
				if e != nil {
					return errors.New(filePath + " gen executor failed:" + e.Error())
				}

				gengoTag := ""
				if len(obj.GengoTags) > preIndex {
					gengoTag = obj.GengoTags[preIndex]
				}
				// run : genexecutor
				e = ioToolkit.RunAttachedCmd("genexecutor", outputFile, gengoTag)
				if e != nil {
					return errors.New(filePath + " " + e.Error() + " . Did you forget to add GOPATH/bin to $PATH environment variable ?")
				}
			}
		}
	}
	return nil
}
