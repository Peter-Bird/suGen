package intern

import "path/filepath"

type Intern struct{}

func (i *Intern) GetDirs(path, name string) []string {
	return []string{
		filepath.Join(path, "internal"),
	}
}

func (i *Intern) GetFiles(path, name string) map[string]string {
	return map[string]string{}
}
