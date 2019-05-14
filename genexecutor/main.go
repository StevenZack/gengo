package main

import (
	"fmt"
	"os"
	"reflect"

	"github.com/StevenZack/tools/ioToolkit"

	"github.com/StevenZack/gengo/gen"

	"xiao/hera/script/init/datastorehouse/roster/verify/rosterverify"
	"github.com/StevenZack/gengo/precompiler/gentag"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("not enough arg")
		return
	}

	s := rosterverify.ClassroomId{}
	t := reflect.TypeOf(s)
	g := &gen.FileGenerator{}
	packageName := "rosterverify"

	fo, e := os.OpenFile(os.Args[1], os.O_TRUNC|os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if e != nil {
		fmt.Println(`create file error :`, e)
		return
	}

	g.Writer = fo
	fo.WriteString("package " + packageName + "\n")

	genGoTag := ""
	if len(os.Args) > 2 {
		genGoTag = os.Args[2]
	}
	str := gentag.Gen(g, genGoTag, t)
	g.WriteAllImports()
	fo.WriteString(str)

	fo.Close()
	ioToolkit.RunAttachedCmd("gofmt", "-w", os.Args[1])
}
