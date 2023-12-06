package vscode

import "path/filepath"

type Vscode struct{}

// GetDirs returns a slice of vscode directory paths
func (vs *Vscode) GetDirs(path string) []string {
	return []string{
		filepath.Join(path, ".vscode"),
	}
}

// GetFiles returns a map of vscode files and functions
func (vs *Vscode) GetFiles(dirName, name string) map[string]string {
	return map[string]string{
		dirName + "/.vscode/settings.json": vs.genSettings(),
		dirName + "/.vscode/launch.json":   vs.genLaunch(),
		dirName + "/.vscode/tasks.json":    vs.genTasks(),
	}
}

/*
-	Understand what to do if theme is not installed
*/
func (vs *Vscode) genSettings() string {
	return `{
	// Editor settings
	"editor.fontSize": 14,
	"editor.lineHeight": 20,
	"editor.tabSize": 4,
	
	// Theme and appearance
	"workbench.colorTheme": "One Dark Pro Monokai Darker",
	"workbench.iconTheme": "material-icon-theme",
	
	// Formatting and linting
	"editor.formatOnSave": true,
	"eslint.alwaysShowStatus": true,
	
	// Language-specific settings
	"[golang]": {
		"editor.insertSpaces": true,
		"editor.tabSize": 4,
		"editor.formatOnSave": true
	},
	
	// Extensions settings
	"liveServer.settings.port": 5500
}`
}

/*
-	Need to address the parameter
-	Todo: Understand launch.json file
-	func (vs *Vscode) GenLaunch(${fileDirname}) string {
*/
func (vs *Vscode) genLaunch() string {
	return `{
	"version": "0.2.0",
	"configurations": [
		{
			"name": "Launch",
			"type": "go",
			"request": "launch",
			"mode": "auto",
			"program": "${fileDirname}",
			"env": {},
			"args": []
		}
	]
}`
}

/*
-	Need to address the parameter
-	Todo: Understand tasks.json file
-	${workspaceFolder}/bin/${fileBasenameNoExtension} ${file}{
*/
func (vs *Vscode) genTasks() string {
	return `{
	"version": "2.0.0",
	"tasks": [
		{
			"label": "Build",
			"type": "shell",
			"command": "go build -o ${workspaceFolder}/bin/${fileBasenameNoExtension} ${file}",
			"group": {
				"kind": "build",
				"isDefault": true
			},
			"problemMatcher": "$go"
		},
		{
			"label": "Run",
			"type": "shell",
			"command": "go run",
			"args": ["${file}"],
			"group": "test",
			"problemMatcher": "$go"
		},
		{
			"label": "Test",
			"type": "shell",
			"command": "go test",
			"args": ["./..."],
			"group": "test",
			"problemMatcher": "$go"
		},
		{
			"label": "Clean",
			"type": "shell",
			"command": "go clean",
			"group": "build",
			"problemMatcher": []
		}
	]
}`
}
