package gen

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/StevenZack/tools/ioToolkit"

	"github.com/StevenZack/tools/strToolkit"

	"github.com/StevenZack/tools/fileToolkit"
)

var verbosely bool

const genExecutorPkgPath = "github.com/StevenZack/gengo/genexecutor"

func SetVerbosely(b bool) {
	verbosely = b
}

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

func compile(pkgPath string) error {
	if !fileToolkit.IsGoPathPkg(pkgPath) {
		return nil
	}

	absPath := fileToolkit.GetGOPATH() + "src/" + strToolkit.Getunpath(pkgPath)
	list, e := fileToolkit.RangeFilesInDir(absPath)
	if e != nil {
		return e
	}

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

		structs, e := ParseFileGengoStructs(filePath)
		if e != nil {
			return errors.New(filePath + " parseFileGengoStructs failed:" + e.Error())
		}

		log("Found", len(structs), "structs", structs)

		for _, obj := range structs {
			outputFile, e := obj.GetGengoFileOutputPath()
			log("output:", outputFile)
			if e != nil {
				return e
			}

			e = os.Remove(outputFile)
			if e != nil {
				log("os.Remove file failed:", outputFile)
			}
			log("os.Remove succeed:", outputFile)

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

func log(args ...interface{}) {
	if verbosely {
		fmt.Println(args...)
	}
}

func generateExecutor(obj GengoStruct) error {
	path := fileToolkit.GetGOPATH() + "src/" + genExecutorPkgPath + "/main.go"
	bakPath := path + ".bak"
	if !fileToolkit.IsFileExists(bakPath) {
		return errors.New("file " + bakPath + " doesn't exists")
	}
	str, e := fileToolkit.ReadFileAll(bakPath)
	if e != nil {
		return errors.New("readFile." + e.Error())
	}
	str = strings.Replace(str, "str := data_gengo.Gen(g, genGoTag, t)", "str := "+obj.PreCompilerPkgName+".Gen(g, genGoTag, t)", -1)
	str = strings.Replace(str, `"github.com/StevenZack/gengo/example/data"`, `"`+obj.StructPkg+`"`, -1)
	structPkgName, e := fileToolkit.GetPkgNameFromPkg(obj.StructPkg)
	if e != nil {
		return errors.New("GetPkgNameFromPkg." + e.Error())
	}
	str = strings.Replace(str, "s := data.Student{}", "s := "+structPkgName+"."+obj.Name+"{}", -1)
	str = strings.Replace(str, `packageName := "data"`, `packageName := "`+structPkgName+`"`, -1)
	str = strings.Replace(str, "github.com/StevenZack/gengo/example/data_gengo", obj.PreCompilerPkg, -1)

	fo, e := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if e != nil {
		return errors.New("os.OpenFile." + e.Error())
	}
	defer fo.Close()
	fo.WriteString(str)

	return ioToolkit.RunAttachedCmd("go", "install", genExecutorPkgPath)
}
