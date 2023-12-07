package docs

import "path/filepath"

type Docs struct{}

func (d *Docs) GetDirs(path, name string) []string {
	return []string{
		filepath.Join(path, "docs"),
		filepath.Join(path, "docs", "HLD"),
		filepath.Join(path, "docs", "PRD"),
	}
}

func (d *Docs) GetFiles(path, name string) map[string]string {
	return map[string]string{}
}
