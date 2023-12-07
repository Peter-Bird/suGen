package main

import (
	"fmt"
	"log"
	"suDir/pkg/app"
	"suDir/pkg/assets"
	"suDir/pkg/compile"
	"suDir/pkg/config"
	"suDir/pkg/docker"
	"suDir/pkg/docs"
	"suDir/pkg/env"
	"suDir/pkg/git"
	"suDir/pkg/github"
	"suDir/pkg/gomod"
	"suDir/pkg/gotree"
	"suDir/pkg/intern"
	"suDir/pkg/license"
	"suDir/pkg/readme"
	"suDir/pkg/scripts"
	"suDir/pkg/shortcuts"
	"suDir/pkg/source"
	"suDir/pkg/tests"
	"suDir/pkg/vendors"
	"suDir/pkg/vscode"

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
		gomod.GoImports(path)
		outputMsg("- Imports Cleared!\n")

		gomod.GoModTidy(path)
		outputMsg("- Modules tidied!\n")

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
