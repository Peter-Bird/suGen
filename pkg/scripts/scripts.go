package scripts

import "path/filepath"

type Scripts struct{}

func (s *Scripts) GetDirs(path, name string) []string {
	return []string{
		filepath.Join(path, "scripts"),
	}
}

func (s *Scripts) GetFiles(path, name string) map[string]string {
	return map[string]string{}
}
