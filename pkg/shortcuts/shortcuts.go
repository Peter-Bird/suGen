package shortcuts

import (
	"fmt"
	"path/filepath"
)

// Shortcuts provides utilities for generating desktop shortcuts.
type Shortcuts struct{}

// GetDirs returns a slice of directory paths related to shortcuts.
func (s *Shortcuts) GetDirs(path, name string) []string {
	return []string{
		filepath.Join(path, "shortcuts"),
	}
}

// GetFiles returns a map of file paths for shortcuts.
func (s *Shortcuts) GetFiles(path, name string) map[string]string {

	file := name + ".desktop"

	return map[string]string{
		filepath.Join(path, "shortcuts", file): s.genDesktop(path, name),
	}
}

// genDesktop generates the content of a linux desktop shortcut file.
func (s *Shortcuts) genDesktop(path, name string) string {
	return fmt.Sprintf(`[Desktop Entry]
Version=1.0
Name=%s
Exec=%s/bin/%s
Path=%s
Icon=%s/assets/%s.png
Type=Application
Categories=Utility;
`,
		name,
		path, name,
		path,
		path, name)
}
