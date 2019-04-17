package gen

import (
	"bufio"
	"errors"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/StevenZack/tools/fileToolkit"
	"github.com/StevenZack/tools/strToolkit"
)

// ParseFileGengoStructs ParseFileGengoStructs
func ParseFileGengoStructs(path string) ([]GengoStruct, error) {
	log("parsing", path)
	f, e := os.OpenFile(path, os.O_RDONLY, 0644)
	if e != nil {
		return nil, e
	}
	defer f.Close()
	r := bufio.NewReader(f)

	structs := []GengoStruct{}
	index := 0
	dir, e := fileToolkit.GetDirOfFile(path)
	if e != nil {
		return nil, e
	}
	structPkg, e := fileToolkit.GetPkgFromDir(dir)
	if e != nil {
		return nil, e
	}
FileLoop:
	for {
		index++
		line, e := fileToolkit.ReadLine(r)
		if e != nil {
			break FileLoop
		}
		if !strings.Contains(line, "gengo ") {
			continue
		}
		precompilers, gengoTags, outputPkgs, e := readGengoFromLine(line)
		if e != nil {
			return nil, e
		}

		gs := GengoStruct{}
		gs.PreCompilers = precompilers
		gs.GengoTags = gengoTags
		gs.OutputPkgs = outputPkgs
		gs.StructPkg = structPkg
		gs.FilePath = path

	GengoLoop:
		for {
			index++
			gengoLine, e := fileToolkit.ReadLine(r)
			if e != nil {
				break FileLoop
			}
			if strings.Contains(gengoLine, "}") {
				break
			}
			if !strings.Contains(gengoLine, "struct {") && !strings.Contains(gengoLine, "struct{") {
				continue
			}
			name, e := readStructNameFromLine(gengoLine)
			if e != nil {
				return nil, errors.New("line " + strconv.Itoa(index) + " : " + e.Error())
			}
			gs.Name = name

			//StructLoop:
			for {
				index++
				structline, e := fileToolkit.ReadLine(r)
				if e != nil {
					break FileLoop
				}
				if strings.Contains(structline, "}") {
					break GengoLoop
				}
			}
		}

		structs = append(structs, gs)
	}

	return structs, nil
}

func readGengoFromLine(l string) ([]PreCompiler, []string, []string, error) {
	formatErr := errors.New("bad gengo format")
	if !strToolkit.StartsWith(l, "//") {
		return nil, nil, nil, formatErr
	}

	index := strings.Index(l, "gengo ")
	if index < 2 {
		return nil, nil, nil, formatErr
	}

	strs := strings.Split(l[index+len("gengo "):], " ")
	if len(strs) == 0 {
		return nil, nil, nil, formatErr
	}
	var precompilers []PreCompiler
	var gengoTags, outputPkgs []string
	for _, v := range strings.Split(strs[0], ",") {
		pkgName, e := fileToolkit.GetPkgNameFromPkg(v)
		if e != nil {
			return nil, nil, nil, errors.New("pkg:" + v + " is not a Go Package")
		}
		pre := PreCompiler{
			Pkg:     v,
			PkgName: pkgName,
		}
		precompilers = append(precompilers, pre)
	}
	if len(strs) > 1 && strs[1] != "" {
		gengoTags = strings.Split(strs[1], ",")
	}
	if len(strs) > 2 && strs[2] != "" {
		outputPkgs = strings.Split(strs[2], ",")
	}

	return precompilers, gengoTags, outputPkgs, nil
}

func readStructNameFromLine(l string) (string, error) {
	l = strings.Replace(l, "\t", "", -1)
	strs := strings.Split(l, " ")
	structIndex := -1
	for index, s := range strs {
		if strToolkit.StartsWith(s, "struct") {
			structIndex = index
			break
		}
	}
	formatErr := errors.New("bad struct format")
	if structIndex == 0 {
		return "", formatErr
	}
	name := strs[structIndex-1]
	if name == "" {
		return "", formatErr
	}
	return name, nil
}

func readFieldsFromLine(l string) ([]Field, error) {
	l = strings.Replace(l, "\t", "", -1)
	formatErr := errors.New("bad struct field format")
	fields := []Field{}
	strs := strings.Split(l, " ")

	if len(strs) < 2 {
		return nil, formatErr
	}

	parts := []string{}
	for _, value := range strs {
		if value != "" {
			parts = append(parts, value)
		}
	}
	if len(parts) < 2 {
		return nil, formatErr
	}
	kind := parts[1]
	if strings.Contains(parts[0], ",") {
		names := strings.Split(parts[0], ",")
		for _, name := range names {
			f := Field{}
			f.Name = name
			f.Kind = kind
			fields = append(fields, f)
		}
		return fields, nil
	}
	f := Field{}
	f.Kind = kind
	f.Name = parts[0]
	if len(parts) > 2 {
		tag := parts[2]
		if !strings.Contains(tag, "`") {
			return nil, formatErr
		}
		f.Tag = reflect.StructTag(tag)
	}
	fields = append(fields, f)
	return fields, nil
}
