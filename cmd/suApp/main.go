package main

import (
	"fmt"
	"log"
	"path/filepath"
	"suApp/pkg/app"
	"suApp/pkg/assets"
	"suApp/pkg/compile"
	"suApp/pkg/config"
	"suApp/pkg/docker"
	"suApp/pkg/docs"
	"suApp/pkg/env"
	"suApp/pkg/git"
	"suApp/pkg/github"
	"suApp/pkg/gomod"
	"suApp/pkg/gotree"
	"suApp/pkg/intern"
	"suApp/pkg/license"
	"suApp/pkg/makefile"
	"suApp/pkg/readme"
	"suApp/pkg/scripts"
	"suApp/pkg/shortcuts"
	"suApp/pkg/source"
	"suApp/pkg/tests"
	"suApp/pkg/vendors"
	"suApp/pkg/vscode"

	"fyne.io/fyne/v2/widget"
)

const (
	ROOT    = "/home/julian/Startups/"
	ID      = "Peter-Bird"
	DIRPERM = 0755
)

// main.go or a root package file
type FileGen interface {
	GetDirs(string, string) []string
	GetFiles(string, string) map[string]string
}

func main() {
	CreateWindow()
}

// Todo: check errors on file creation
func appGen(fileGen FileGen, path, name string) {

	dirs := fileGen.GetDirs(path, name)
	err := gotree.CreateDirs(dirs, DIRPERM, outputMsg)
	if err != nil {
		log.Fatalf("Failed to create directories: %v", err)
	}

	files := fileGen.GetFiles(path, name)
	gotree.CreateFiles(files, outputMsg)
}

func buildApp(name string, output *widget.Entry) error {
	outputMsg(fmt.Sprintln("Program: suApp Started"))

	path := ROOT + name

	if err := gotree.CheckDirExists(path); err != nil {
		log.Fatalf("Directory exists: %s", err)
	}

	selected := getSelectedCheckboxes(checkboxes)
	actions := getActions(path, name)
	execChecked(selected, actions)

	if Contains(selected, "mod") {
		// Added dir for GoImports after moving main.go to cmd
		dir := filepath.Join(path, "cmd", name)
		gomod.GoImports(dir)
		outputMsg("- Imports Cleared!\n")

		gomod.GoModTidy(dir)
		outputMsg("- Modules Tidied!\n")

		if Contains(selected, "vendor") {
			gomod.GoModVendor(path)
			outputMsg("- Modules Vendored!\n")
		}
	}

	if Contains(selected, "git") {
		if err := git.InitRepo(path); err != nil {
			log.Fatalf("Error initializing git: %s", err)
		}
		outputMsg("- Git Initialized!\n")
	}

	if Contains(selected, "bin") {
		compile.Compiler(path, name, outputMsg)
		outputMsg("- Application Compiled!\n\n")
	}

	outputMsg("Completed Succesfully")

	return nil
}

// Contains checks if a string is present in a slice.
func Contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}

func getActions(path, name string) map[string]func() {
	types := map[string]interface{}{
		"app":       &app.App{},
		"mod":       &gomod.Gomod{},
		"bin":       &compile.Compile{},
		"source":    &source.Source{},
		"git":       &git.Git{},
		"config":    &config.Config{},
		"env":       &env.Env{},
		"vscode":    &vscode.Vscode{},
		"github":    &github.Github{},
		"makefile":  &makefile.Makefile{},
		"docker":    &docker.Docker{},
		"documents": &docs.Docs{},
		"shortcuts": &shortcuts.Shortcuts{},
		"assets":    &assets.Assets{},
		"tests":     &tests.Tests{},
		"internal":  &intern.Intern{},
		"scripts":   &scripts.Scripts{},
		"license":   &license.License{},
		"readme":    &readme.Readme{},
		"vendor":    &vendors.Vendors{},
	}

	actions := make(map[string]func())
	for key, val := range types {
		fileGenVal, ok := val.(FileGen)
		if !ok {
			log.Fatalf("Type for key %s does not implement FileGen interface", key)
		}

		actions[key] = func() {
			appGen(fileGenVal, path, name)
		}
	}

	return actions
}

func execChecked(selected []string, funcMap map[string]func()) {
	for _, checkbox := range selected {
		if funcToExecute, exists := funcMap[checkbox]; exists {
			funcToExecute()
		}
	}
}
