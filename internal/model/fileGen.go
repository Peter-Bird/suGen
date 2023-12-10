package model

// main.go or a root package file
type FileGen interface {
	GetDirs(string, string) []string
	GetFiles(string, string) map[string]string
}
