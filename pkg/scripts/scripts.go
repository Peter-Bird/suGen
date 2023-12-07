package scripts

import "path/filepath"

type Scripts struct{}

func (s *Scripts) GetDirs(path string) []string {
	return []string{
		filepath.Join(path, "scripts"),
	}
}

func (s *Scripts) GetFiles(dirName, name string) map[string]string {
	return map[string]string{}
}
