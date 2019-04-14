package main

import (
	"flag"

	"github.com/StevenZack/gengo/gen"

	"github.com/StevenZack/gengo/help"
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

	switch action {
	case "gen":
		gen.Gen(args)
	}
}
