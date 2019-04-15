package main

import (
	"flag"
	"fmt"

	"github.com/StevenZack/gengo/help"
	"github.com/StevenZack/tools/ioToolkit"

	"github.com/StevenZack/gengo/gen"
)

var verbose = flag.Bool("v", false, "show log")

func main() {
	flag.Parse()
	gen.SetVerbosely(*verbose)
	if len(flag.Args()) < 1 {
		help.ShowAll()
		return
	}

	action := flag.Arg(0)

	var args []string
	if len(flag.Args()) == 1 {
		args = nil
	} else {
		args = flag.Args()[1:]
	}
	if *verbose {
		fmt.Println("action =", action)
	}
	switch action {
	case "gen":
		gen.Gen(args)
	case "install":
		gen.Gen(args)
		runGoCmd(action, args)
	case "build":
		gen.Gen(args)
		runGoCmd(action, args)
	default:
		runGoCmd(action, args)
	}
}
func runGoCmd(action string, args []string) {
	e := ioToolkit.RunAttachedCmd("go", append([]string{action}, args...)...)
	if e != nil {
		fmt.Println("go "+action+" error :", e)
		return
	}
}
