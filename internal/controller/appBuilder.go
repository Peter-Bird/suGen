package controller

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

	"suApp/internal/model"
	"suApp/internal/view"

	"fyne.io/fyne/v2/widget"
)

const (
	ROOT    = "/home/julian/Startups/"
	DIRPERM = 0755
)

type Cntl struct {
	gui view.View
}

// Todo: check errors on file creation
func (c *Cntl) appGen(fileGen model.FileGen, path, name string) {

	dirs := fileGen.GetDirs(path, name)
	err := gotree.CreateDirs(dirs, DIRPERM, c.gui.OutputMsg)
	if err != nil {
		log.Fatalf("Failed to create directories: %v", err)
	}

	files := fileGen.GetFiles(path, name)
	gotree.CreateFiles(files, c.gui.OutputMsg)
}

func (c *Cntl) BuildApp(name string, output *widget.Entry) error {
	c.gui.OutputMsg(fmt.Sprintln("Program: suApp Started"))

	path := ROOT + name

	if err := gotree.CheckDirExists(path); err != nil {
		log.Fatalf("Directory exists: %s", err)
	}

	selected := c.gui.GetSelectedCheckboxes(view.Checkboxes)
	actions := c.getActions(path, name)
	c.execChecked(selected, actions)

	if c.Contains(selected, "mod") {
		// Added dir for GoImports after moving main.go to cmd
		dir := filepath.Join(path, "cmd", name)
		gomod.GoImports(dir)
		c.gui.OutputMsg("- Imports Cleared!\n")

		gomod.GoModTidy(dir)
		c.gui.OutputMsg("- Modules Tidied!\n")

		if c.Contains(selected, "vendor") {
			gomod.GoModVendor(path)
			c.gui.OutputMsg("- Modules Vendored!\n")
		}
	}

	if c.Contains(selected, "git") {
		if err := git.InitRepo(path); err != nil {
			log.Fatalf("Error initializing git: %s", err)
		}
		c.gui.OutputMsg("- Git Initialized!\n")
	}

	if c.Contains(selected, "bin") {
		compile.Compiler(path, name, c.gui.OutputMsg)
		c.gui.OutputMsg("- Application Compiled!\n\n")
	}

	c.gui.OutputMsg("Completed Succesfully")

	return nil
}

// Contains checks if a string is present in a slice.
func (c *Cntl) Contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}

func (c *Cntl) getActions(path, name string) map[string]func() {
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
		fileGenVal, ok := val.(model.FileGen)
		if !ok {
			log.Fatalf("Type for key %s does not implement FileGen interface", key)
		}

		actions[key] = func() {
			c.appGen(fileGenVal, path, name)
		}
	}

	return actions
}

func (c *Cntl) execChecked(selected []string, funcMap map[string]func()) {
	for _, checkbox := range selected {
		if funcToExecute, exists := funcMap[checkbox]; exists {
			funcToExecute()
		}
	}
}
