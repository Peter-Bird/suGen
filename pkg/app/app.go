package app

import "path/filepath"

type App struct{}

func (a *App) GetDirs(path, name string) []string {
	return []string{
		filepath.Join(path, ""),
	}
}

func (a *App) GetFiles(path, name string) map[string]string {
	return map[string]string{}
}
