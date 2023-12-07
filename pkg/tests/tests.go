package tests

import "path/filepath"

type Tests struct{}

func (t *Tests) GetDirs(path, name string) []string {
	return []string{
		filepath.Join(path, "test"),
	}
}

func (t *Tests) GetFiles(path, name string) map[string]string {
	return map[string]string{}
}
