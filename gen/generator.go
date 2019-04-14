package gen

import "os"

// FileGenerator generate _gengo.go files
type FileGenerator struct {
	Writer *os.File
}

func (f *FileGenerator) AddImport(s string) error {
	_, e := f.Writer.WriteString("import \"" + s + "\"\n")
	if e != nil {
		return e
	}
	return nil
}
