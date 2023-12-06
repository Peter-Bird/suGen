package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var checkboxes []*widget.Check
var outputEntry *widget.Entry

func CreateWindow() {
	myApp := app.NewWithID(ID)
	myWin := myApp.NewWindow("Create Application")

	// Label for Application Name
	appNameLabel := widget.NewLabel("Name")

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Enter the application's Name")

	// Label for Application Name
	gridLabel := widget.NewLabel("Integrations")

	// Create an array (slice) of checkboxes
	checkboxes = []*widget.Check{
		widget.NewCheck("app", func(checked bool) {}),
		widget.NewCheck("mod", func(checked bool) {}),
		widget.NewCheck("source", func(checked bool) {}),
		widget.NewCheck("bin", func(checked bool) {}),
		widget.NewCheck("git", func(checked bool) {}),
		widget.NewCheck("vscode", func(checked bool) {}),
		widget.NewCheck("github", func(checked bool) {}),
		widget.NewCheck("docker", func(checked bool) {}),
		widget.NewCheck("shortcuts", func(checked bool) {}),
		widget.NewCheck("documents", func(checked bool) {}),
		widget.NewCheck("scripts", func(checked bool) {}),
		widget.NewCheck("tests", func(checked bool) {}),
		widget.NewCheck("config", func(checked bool) {}),
		widget.NewCheck("assets", func(checked bool) {}),
		widget.NewCheck("env", func(checked bool) {}),
		widget.NewCheck("readme", func(checked bool) {}),
		widget.NewCheck("license", func(checked bool) {}),
		widget.NewCheck("internal", func(checked bool) {}),
		widget.NewCheck("vendor", func(checked bool) {}),
	}

	// Adding checkboxes to a grid
	checkboxGrid := container.NewGridWithColumns(4)
	for i, checkbox := range checkboxes {
		if i < 4 {
			checkbox.Checked = true
		}
		checkboxGrid.Add(checkbox)
	}

	outputEntry = widget.NewMultiLineEntry()
	outputEntry.SetPlaceHolder("Output will be shown here")
	outputEntry.SetMinRowsVisible(16)

	startButton := widget.NewButton("Generate", func() {
		go buildApp(nameEntry.Text, outputEntry)
	})

	// Setting up a key listener for the nameEntry
	nameEntry.OnSubmitted = func(s string) {
		startButton.OnTapped() // Trigger the button's action
	}

	myWin.SetContent(container.NewVBox(
		appNameLabel,
		nameEntry,
		gridLabel,
		checkboxGrid,
		startButton,
		outputEntry,
	))

	myWin.Resize(fyne.NewSize(600, 420))
	myWin.ShowAndRun()
}

func getSelectedCheckboxes(checkboxes []*widget.Check) []string {
	var selected []string
	for _, checkbox := range checkboxes {
		if checkbox.Checked {
			selected = append(selected, checkbox.Text)
		}
	}
	return selected
}

func outputMsg(txt string) {
	outputEntry.Append(txt)
}
