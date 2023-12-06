package config

import "path/filepath"

type Config struct{}

func (c *Config) GetDirs(path string) []string {
	return []string{
		filepath.Join(path, "config"),
	}
}

func (c *Config) GetFiles(path, appName string) map[string]string {

	return map[string]string{
		filepath.Join(path, "config", ".env.example"): c.genDotEnvExample(),
	}
}

func (c *Config) genDotEnvExample() string {
	return `DB_HOST=localhost
DB_USER=user
DB_PASSWORD=password
API_KEY=your-api-key
`
}
