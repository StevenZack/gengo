package gen

import (
	"errors"
	"os"
	"strings"

	"github.com/StevenZack/tools/fileToolkit"
	"github.com/StevenZack/tools/ioToolkit"
)

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
	str = strings.Replace(str, "github.com/StevenZack/gengo/example/precompiler/tostring_gengo", obj.PreCompilerPkg, -1)

	fo, e := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if e != nil {
		return errors.New("os.OpenFile." + e.Error())
	}
	defer fo.Close()
	fo.WriteString(str)

	return ioToolkit.RunAttachedCmd("go", "install", genExecutorPkgPath)
}
