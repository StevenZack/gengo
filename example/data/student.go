package data

import (
	"github.com/StevenZack/gengo/example/base"
)

// Student stores a student . gengo github.com/StevenZack/gengo/example/precompiler/tostring_gengo gengo_tag model/model.go
type Student struct {
	base.Base
	Name   string `json:"name"`
	Age    int
	School string
}
