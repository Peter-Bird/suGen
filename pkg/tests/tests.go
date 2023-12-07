package tests

import "path/filepath"

type Tests struct{}

func (t *Tests) GetDirs(path string) []string {
	return []string{
		filepath.Join(path, "test"),
	}
}

func (t *Tests) GetFiles(dirName, name string) map[string]string {
	return map[string]string{}
}
