package docs

import "path/filepath"

type Docs struct{}

func (d *Docs) GetDirs(path string) []string {
	return []string{
		filepath.Join(path, "docs"),
		filepath.Join(path, "docs", "HLD"),
		filepath.Join(path, "docs", "PRD"),
	}
}

func (d *Docs) GetFiles(dirName, name string) map[string]string {
	return map[string]string{}
}
