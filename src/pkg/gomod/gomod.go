package gomod

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

type Gomod struct{}

func (m *Gomod) GetDirs(path string) []string {
	return []string{}
}

func (m *Gomod) GetFiles(path, name string) map[string]string {
	return map[string]string{
		path + "/go.mod": m.GenGomod(name),
	}
}

func GoImports(path string) {
	cmd := exec.Command("goimports", "-w", path+"/main.go")
	if err := cmd.Run(); err != nil {
		log.Fatalf("Error cleaning imports: %s", err)
	}
}

func GoGetModule(module string) error {
	getCmd := exec.Command("go", "get", module)
	return getCmd.Run()
}

func GoModTidy(dirName string) {
	cmd := exec.Command("go", "mod", "tidy", "-e")
	cmd.Dir = dirName
	if err := cmd.Run(); err != nil {
		log.Fatalf("Error Tidying: %s", err)
	}
}

func GoModVendor(dirName string) {

	var stdout, stderr bytes.Buffer

	cmd := exec.Command("go", "mod", "vendor")
	cmd.Dir = dirName
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		log.Printf("stdout: %s", stdout.String())
		log.Printf("stderr: %s", stderr.String())
		log.Fatalf("Error Vendoring: %s", err)
	}
}

func (m *Gomod) GenGomod(name string) string {

	mod := fmt.Sprintf(`module peter-bird.com/%s

go 1.21.1

replace peter-bird.com/apikey => /home/julian/go/pkgs/apikey

replace peter-bird.com/chatgpt => /home/julian/go/pkgs/chatgpt
`, name)

	return mod
}
