package tool

import (
	"bufio"
	"errors"
	"os"
	"strings"

	"github.com/StevenZack/tools/fileToolkit"
)

func GetFirstGoFile(dir string) (string, error) {

	files, e := fileToolkit.RangeFilesInDir(dir)
	if e != nil {
		return "", e
	}
	if len(files) == 0 {
		return "", errors.New("no .go files in dir:" + dir)
	}
	return files[0], nil
}

func ReadPkgFromLine(l string) (string, error) {
	l = strings.Replace(l, "\t", "", -1)
	strs := strings.Split(l, " ")
	if len(strs) < 2 {
		return "", errors.New("bad package format")
	}
	return strs[1], nil
}

func ReadFirstLine(filePath string) (string, error) {
	f, e := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if e != nil {
		return "", e
	}
	defer f.Close()
	r := bufio.NewReader(f)
	line, e := ReadLine(r)
	return line, e
}

func ReadLine(r *bufio.Reader) (string, error) {
	b, _, e := r.ReadLine()
	if e != nil {
		return "", e
	}
	return string(b), nil
}
