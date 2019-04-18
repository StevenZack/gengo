package gentag

import (
	"reflect"

	"github.com/StevenZack/tools/strToolkit"

	"github.com/StevenZack/gengo/gen"
)

func Gen(g *gen.FileGenerator, gengoTag string, t reflect.Type) string {
	str := `type ` + t.Name() + ` struct{
`
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		snake := strToolkit.ToSnakeCase(field.Name)
		str += field.Name + " " + field.Type.Name() + " `json:\"" + snake + "\" db:\"" + snake + "\"`\n"
	}
	str += "}\n"
	return str
}
