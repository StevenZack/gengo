package tool

import "testing"

func Test_getImportFromL(t *testing.T) {
	t.Log(GetAllImports("/Users/stevenzacker/go/src/github.com/StevenZack/gengo/tool/import.go"))
}
