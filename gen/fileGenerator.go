package gen

import "os"

// FileGenerator generate _gengo.go files
type FileGenerator struct {
	Writer  *os.File
	imports map[string]bool
}

// AddImport adds import s into output go file
func (f *FileGenerator) AddImport(s string) {
	if f.imports == nil {
		f.imports = make(map[string]bool)
	}
	_, ok := f.imports[s]
	if ok {
		return
	}

	f.imports[s] = true
	return
}

// RemoveImport removes import s if exists
func (f *FileGenerator) RemoveImport(s string) {
	if f.imports == nil {
		return
	}
	_, ok := f.imports[s]
	if !ok {
		return
	}
	delete(f.imports, s)
}

// WriteAllImports write all import-operation you did into output go file
func (f *FileGenerator) WriteAllImports() error {
	if f.imports == nil {
		return nil
	}
	for key := range f.imports {
		_, e := f.Writer.WriteString("import \"" + key + "\"\n")
		if e != nil {
			return e
		}
	}
	return nil
}
