package tool

import (
	"bufio"
	"errors"
	"os"
	"strings"

	"github.com/StevenZack/tools/strToolkit"
)

func GetAllImports(path string) ([]string, error) {
	imports := []string{}

	f, e := os.Open(path)
	if e != nil {
		return nil, e
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, e := readLine(r)
		if e != nil {
			break
		}
		if !strToolkit.StartsWith(line, "import") {
			continue
		}

		if strings.Contains(line, "(") {
			for {
				l, e := readLine(r)
				if e != nil {
					return nil, e
				}
				if strings.Contains(l, ")") {
					break
				}
				imp, e := getImportFromL(l)
				if e != nil {
					continue
				}
				imports = append(imports, imp)
			}
			continue
		}

		imp, e := getImportFromL(line)
		if e != nil {
			continue
		}
		imports = append(imports, imp)
	}

	return imports, nil
}

func readLine(r *bufio.Reader) (string, error) {
	b, _, e := r.ReadLine()
	if e != nil {
		return "", e
	}
	return string(b), nil
}

func getImportFromL(l string) (string, error) {
	list := strings.Split(l, " ")
	for _, str := range list {
		str = strings.Replace(str, "\t", "", -1)
		count := strings.Count(str, `"`)
		if count != 2 {
			continue
		}
		if !strToolkit.StartsWith(str, `"`) {
			continue
		}
		if !strToolkit.EndsWith(str, `"`) {
			continue
		}
		return str[1 : len(str)-1], nil
	}
	return "", errors.New("not found")
}

