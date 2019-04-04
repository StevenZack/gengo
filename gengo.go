package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/StevenZack/tools/fileToolkit"
	"github.com/StevenZack/tools/strToolkit"
)

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		gengoPackage("./")
		return
	}
	target := "."
	if len(flag.Args()) == 1 {
		parseCommand(target)
		return
	}
	target = flag.Arg(1)
	gengoTarget(target)
}
func gengoTarget(target string) {
	ok, e := fileToolkit.IsDir(target)
	if e != nil {
		fmt.Println(` error :`, e)
		return
	}
	if ok {
		gengoPackage(target)
		return
	}
	gengoFile(target)
}
func gengoPackage(dir string) {
	e := rangeGenGoFilesInDir(dir, func(path string) {
		gengoFile(path)
	})
	if e != nil {
		fmt.Println(`range dir error :`, e)
		return
	}
}
func gengoFile(path string) {
	ok, e := fileToolkit.IsDir(path)
	if e != nil {
		fmt.Println(` error :`, e)
		return
	}
	if ok {
		fmt.Println("path:", path, "is not file")
		return
	}
	if !strToolkit.EndsWith(path, ".gengo") {
		fmt.Println("path:", path, "is not gengo file")
		return
	}

}
func rangeGenGoFilesInDir(dir string, fn func(path string)) error {
	infos, e := ioutil.ReadDir(dir)
	if e != nil {
		return e
	}
	for _, info := range infos {
		if strToolkit.EndsWith(info.Name(), ".gengo") {
			fn(fileToolkit.Getrpath(dir) + info.Name())
		}
	}
	return nil
}

func parseCommand(target string) {
	command := flag.Arg(0)
	switch command {
	case "gen":
		gengoTarget(target)
	case "run":
		gengoTarget(target)
		cmdToolkit.
	}
}
