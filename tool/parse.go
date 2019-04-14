package tool

import "os"

func ParseFile(path string) error {
	f, e := os.OpenFile(path, os.O_RDONLY, 0644)
	if e != nil {
		return e
	}
	defer f.Close()
	
}
