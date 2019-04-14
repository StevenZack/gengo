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

	f, e := os.OpenFile(path, os.O_RDONLY, 0644)
	if e != nil {
		return nil, e
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, e := readLine(r)
		if e != nil {
			return nil, e
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

			}
		}
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
