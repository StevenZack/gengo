package json_gengo

import (
	"reflect"

	"github.com/StevenZack/gengo/gen"
)

func Gen(g *gen.FileGenerator, gengoTag string, t reflect.Type) string {
	str := `func (s *` + t.Name() + `)JSONObject()string{
	b,e:=json.Marshal(s)
	if e!=nil{
		return "{}"
	}
	return string(b)
}`
	g.AddImport("encoding/json")
	return str
}
