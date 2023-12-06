package git

import (
	"fmt"
	"os/exec"
)

type Git struct{}

func (g *Git) GetDirs(path string) []string {
	return []string{}
}

func (g *Git) GetFiles(path, name string) map[string]string {
	return map[string]string{
		path + "/.gitignore": g.genGitIgnore(),
	}
}

// GitCommandFunc type for executing git commands
type GitCommandFunc func(dirName, command string, args ...string) error

// gitCommands map to associate git commands with their respective function
var gitCommands = map[string]GitCommandFunc{
	"init":   executeGitCommand,
	"add":    executeGitCommand,
	"commit": executeGitCommand,
}

// executeGitCommand executes a git command
func executeGitCommand(dirName, command string, args ...string) error {
	cmd := exec.Command("git", append([]string{command}, args...)...)
	cmd.Dir = dirName
	return cmd.Run()
}

// executeGit executes a git command based on the command name and arguments
func executeGit(dirName, command string, args ...string) error {
	if gitFunc, exists := gitCommands[command]; exists {
		return gitFunc(dirName, command, args...)
	}
	return fmt.Errorf("unknown git command: %s", command)
}

// initializeGitRepository initializes a git repository with commit
func InitRepo(dirName string) error {
	if err := executeGit(dirName, "init"); err != nil {
		return fmt.Errorf("error initializing git repository with 'init': %s", err)
	}
	if err := executeGit(dirName, "add", "."); err != nil {
		return fmt.Errorf("error adding files to git with 'add': %s", err)
	}
	if err := executeGit(dirName, "commit", "-m", "Initial Load"); err != nil {
		return fmt.Errorf("error committing to git with 'commit': %s", err)
	}
	return nil
}

func (g *Git) genGitIgnore() string {
	return `# Ignore everything
*

# But not these files...
!/.gitignore

!*.go
!go.sum
!go.mod

!README.md
!LICENSE

# !Makefile

# ...even if they are in subdirectories
!*/
`
}
