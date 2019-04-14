package tool

import "testing"

func Test_getImportFromL(t *testing.T) {
	t.Log(ParseFileGengoStructs("/Users/stevenzacker/go/src/github.com/StevenZack/gengo/example/data/student.go"))
}
