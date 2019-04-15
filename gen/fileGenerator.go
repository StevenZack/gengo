package gen

import "os"

// FileGenerator generate _gengo.go files
type FileGenerator struct {
	Writer  *os.File
	imports map[string]bool
}

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

func (f *FileGenerator) WriteAllImports() error {
	if f.imports == nil {
		return nil
	}
	for key, _ := range f.imports {
		_, e := f.Writer.WriteString("import \"" + key + "\"\n")
		if e != nil {
			return e
		}
	}
	return nil
}
