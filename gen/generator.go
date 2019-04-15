package gen

import "os"

// FileGenerator generate _gengo.go files
type FileGenerator struct {
	Writer  *os.File
	imports map[string]bool
}

func (f *FileGenerator) AddImport(s string) error {
	// if f.imports == nil {
	// 	f.imports = make(map[string]bool)
	// }
	// _, ok := f.imports[s]
	// if ok {
	// 	return nil
	// }

	_, e := f.Writer.WriteString("import \"" + s + "\"\n")
	if e != nil {
		return e
	}
	// f.imports[s] = true
	return nil
}
