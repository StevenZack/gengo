package main

import (
	"os"

	"github.com/StevenZack/gengo/gen"

	"github.com/StevenZack/gengo/help"
)

func main() {
	if len(os.Args) < 2 {
		help.ShowAll()
		return
	}
	action := os.Args[1]

	var args []string
	if len(os.Args) == 2 {
		args = nil
	} else {
		args = os.Args[2:]
	}

	switch action {
	case "gen":
		gen.Gen(args)
	}
}
