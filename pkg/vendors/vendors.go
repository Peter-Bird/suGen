package vendors

import (
	"path/filepath"
	"suApp/pkg/gomod"
)

type Vendors struct{}

// GetDirs returns a slice of vendor directory paths
func (v *Vendors) GetDirs(path, name string) []string {
	return []string{
		filepath.Join(path, "vendor"),
	}
}

// GetFiles activates the mod vendor function
// returns an empty slice to meet interface structure
func (v *Vendors) GetFiles(path, name string) map[string]string {

	gomod.GoModVendor(path)

	return map[string]string{}
}
