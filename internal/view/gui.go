package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const (
	ID = "Peter-Bird"
)

type View struct{}

var Checkboxes []*widget.Check
var outputEntry *widget.Entry

func (v *View) CreateWindow(callback func(string, *widget.Entry) error) {
	myApp := app.NewWithID(ID)
	myWin := myApp.NewWindow("Create Application")

	// Label for Application Name
	appNameLabel := widget.NewLabel("Name")

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Enter the application's Name")

	// Label for Application Name
	gridLabel := widget.NewLabel("Aspects")

	// Define the labels for the checkboxes
	labels := []string{
		"app", "mod", "source", "bin", "git", "makefile",
		"config", "env", "vscode", "github", "docker",
		"shortcuts", "documents", "scripts", "tests",
		"assets", "readme", "license", "internal", "vendor",
	}

	// Adding checkboxes to a grid
	checkboxGrid := container.NewGridWithColumns(4)

	// Loop through the labels and create a checkbox for each
	for i, label := range labels {
		checkbox := widget.NewCheck(label, func(checked bool) {})
		if i < 4 {
			checkbox.Checked = true
		}
		Checkboxes = append(Checkboxes, checkbox)
		checkboxGrid.Add(checkbox)
	}

	outputEntry = widget.NewMultiLineEntry()
	outputEntry.SetPlaceHolder("Output will be shown here")
	outputEntry.SetMinRowsVisible(16)

	startButton := widget.NewButton("Generate", func() {
		go callback(nameEntry.Text, outputEntry)
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

func (v *View) GetSelectedCheckboxes(checkboxes []*widget.Check) []string {
	var selected []string
	for _, checkbox := range checkboxes {
		if checkbox.Checked {
			selected = append(selected, checkbox.Text)
		}
	}
	return selected
}

func (v *View) OutputMsg(txt string) {
	outputEntry.Append(txt)
}
