package main

import (
	"fmt"
	"log"
	"suDir/src/pkg/assets"
	"suDir/src/pkg/compile"
	"suDir/src/pkg/config"
	"suDir/src/pkg/docker"
	"suDir/src/pkg/docs"
	"suDir/src/pkg/git"
	"suDir/src/pkg/github"
	"suDir/src/pkg/gomod"
	"suDir/src/pkg/gotree"
	"suDir/src/pkg/intern"
	"suDir/src/pkg/license"
	"suDir/src/pkg/scripts"
	"suDir/src/pkg/shortcuts"
	"suDir/src/pkg/source"
	"suDir/src/pkg/tests"
	"suDir/src/pkg/vendors"
	"suDir/src/pkg/vscode"

	"fyne.io/fyne/v2/widget"
)

const (
	ROOT    = "/home/julian/Startups/"
	ID      = "Peter-Bird"
	DIRPERM = 0755
)

// main.go or a root package file
type FileGen interface {
	GetDirs(string) []string
	GetFiles(string, string) map[string]string
}

func main() {
	CreateWindow()
}

// Todo: check errors on file creation
func appGen(fileGen FileGen, path, name string) {

	dirs := fileGen.GetDirs(path)
	err := gotree.CreateDirs(dirs, DIRPERM, outputMsg)
	if err != nil {
		log.Fatalf("Failed to create directories: %v", err)
	}

	files := fileGen.GetFiles(path, name)
	createFiles(files)
}

func buildApp(name string, output *widget.Entry) error {
	outputMsg(fmt.Sprintln("Program: suApp Started"))

	path := ROOT + name

	// Check if the directory exists
	if err := gotree.CheckDirExists(path); err != nil {
		log.Fatalf("Directory exists: %s", err)
	}

	if err := setupDirectories(path); err != nil {
		return fmt.Errorf("failed to set up directories: %w", err)
	}

	filesToCreate := gotree.GetFilesToCreate(path, name)
	createFiles(filesToCreate)

	gomod.GoImports(path)
	outputMsg("- Imports Cleared!\n")

	gomod.GoModTidy(path)
	outputMsg("- Modules tidied!\n")

	// gomod.GoModVendor(path)
	// outputMsg("- Modules Vendored!\n")

	if err := git.InitRepo(path); err != nil {
		log.Fatalf("Error initializing git: %s", err)
	}
	outputMsg("- Git Initialized!\n")

	compile.Compiler(path, name, outputMsg)
	outputMsg("- Application Compiled!\n\n")

	selected := getSelectedCheckboxes(checkboxes)
	actions := getActions(path, name)
	execChecked(selected, actions)

	outputMsg("Completed Succesfully")

	return nil
}

func setupDirectories(path string) error {
	dirs := gotree.DefaultDirectories(path)
	return gotree.CreateDirs(dirs, DIRPERM, outputMsg)
}

func getActions(path, name string) map[string]func() {
	return map[string]func(){
		"vscode":    func() { appGen(&vscode.Vscode{}, path, name) },
		"git":       func() { appGen(&git.Git{}, path, name) },
		"github":    func() { appGen(&github.Github{}, path, name) },
		"vendor":    func() { appGen(&vendors.Vendors{}, path, name) },
		"docker":    func() { appGen(&docker.Docker{}, path, name) },
		"documents": func() { appGen(&docs.Docs{}, path, name) },
		"shortcuts": func() { appGen(&shortcuts.Shortcuts{}, path, name) },
		"assets":    func() { appGen(&assets.Assets{}, path, name) },
		"config":    func() { appGen(&config.Config{}, path, name) },
		"tests":     func() { appGen(&tests.Tests{}, path, name) },
		"internal":  func() { appGen(&intern.Intern{}, path, name) },
		"scripts":   func() { appGen(&scripts.Scripts{}, path, name) },
		"source":    func() { appGen(&source.Source{}, path, name) },
		"license":   func() { appGen(&license.License{}, path, name) },
		"gomod":     func() { appGen(&gomod.Gomod{}, path, name) },
		"bin":       func() { appGen(&compile.Compile{}, path, name) },
	}
}

func createFiles(files map[string]string) {
	for fileName, content := range files {
		if err := gotree.MakeFile(fileName, content, outputMsg); err != nil {
			log.Fatalf("Error creating or writing to file %s: %s", fileName, err)
		}
	}
}

// Step 3: Function to Execute Relevant Functions for Checked Checkboxes
func execChecked(selected []string, funcMap map[string]func()) {
	for _, checkbox := range selected {
		if funcToExecute, exists := funcMap[checkbox]; exists {
			funcToExecute()
		}
	}
}
