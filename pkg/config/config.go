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
		filepath.Join(path, "config", "config.json"): c.genConfig(),
	}
}

func (c *Config) genConfig() string {
	return `{}
`
}
