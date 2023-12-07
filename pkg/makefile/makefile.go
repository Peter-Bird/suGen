package makefile

import "fmt"

type Makefile struct{}

func (m *Makefile) GetDirs(path, name string) []string {
	return []string{}
}

func (m *Makefile) GetFiles(path, name string) map[string]string {
	return map[string]string{
		path + "/Makefile": m.genMakeFile(path, name),
	}
}

func (m *Makefile) genMakeFile(path, name string) string {
	return fmt.Sprintf(`build:
	go build -o %s/bin/%s ./cmd/%s
`, path, name, name)
}
