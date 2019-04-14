package tool

import "testing"

func Test_getImportFromL(t *testing.T) {
	t.Log(getImportFromL(`qwe "github.com/stevenzack/qwe"`))
}
