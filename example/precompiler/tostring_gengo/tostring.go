package tostring_gengo

import (
	"reflect"
	"strings"

	"github.com/StevenZack/gengo/gen"
)

func Gen(g *gen.FileGenerator, gengoTag string, t reflect.Type) string {

	str := `func (s *` + t.Name() + `) ToString () string {
	return `

	for index := 0; index < t.NumField(); index++ {
		field := t.Field(index)
		kind := field.Type.Kind().String()
		switch kind {
		case "string":
			str += "s." + field.Name + "+"
		default:
			str += "fmt.Sprint(s." + field.Name + ")+"
			g.AddImport("fmt")
		}
	}

	str = strings.TrimSuffix(str, "+")

	str += "\n}"

	return str
}
