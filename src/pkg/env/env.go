package env

import "path/filepath"

type Env struct{}

func (e *Env) GetDirs(path string) []string {
	return []string{
		filepath.Join(path, ".env"),
	}
}

func (e *Env) GetFiles(path, appName string) map[string]string {

	return map[string]string{
		filepath.Join(path, ".env.example"): e.genDotEnvExample(),
	}
}

func (e *Env) genDotEnvExample() string {
	return `DB_HOST=localhost
DB_USER=user
DB_PASSWORD=password
API_KEY=your-api-key
`
}
