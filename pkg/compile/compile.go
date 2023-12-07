package compile

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"suApp/pkg/gomod"
)

type Compile struct{}

func (c *Compile) GetDirs(path string) []string {
	return []string{
		filepath.Join(path, "bin"),
	}
}

func (c *Compile) GetFiles(path, name string) map[string]string {
	return map[string]string{}
}

func Compiler(path, name string, logger func(string)) error {

	outputPath := path + "/bin/" + name
	cmd := exec.Command("go", "build", "-o", outputPath, path)
	cmd.Dir = path

	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return analyzeBuildError(stderr.String())
	}

	logger(fmt.Sprintf("- Application Compiled: %s\n", outputPath))
	return nil
}

func analyzeBuildError(stderr string) error {
	scanner := bufio.NewScanner(strings.NewReader(stderr))
	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile(`no required module provides package ([^;]+)`)
		matches := re.FindStringSubmatch(line)
		if len(matches) > 1 {
			missingModule := matches[1]
			if err := gomod.GoGetModule(missingModule); err != nil {
				return fmt.Errorf("failed to install missing module %s: %w", missingModule, err)
			}
		}
	}
	return scanner.Err()
}
