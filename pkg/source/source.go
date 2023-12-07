package source

import "path/filepath"

type Source struct{}

// GetDirs returns a slice of source directory paths
func (s *Source) GetDirs(path string) []string {
	return []string{
		filepath.Join(path, "cmd"),
		filepath.Join(path, "pkg"),
	}
}

// GetFiles returns a map of source files and functions
func (s *Source) GetFiles(path, name string) map[string]string {
	return map[string]string{
		path + "/main.go": s.GenMain(),
		path + "/gui.go":  s.GenGui(),
	}
}

func (s *Source) GenMain() string {

	temp := `package main

	import (
		"fmt"
		"os"
		"path/filepath"
	
		"fyne.io/fyne/v2/app"
	
		apiKey "peter-bird.com/apikey"
		chatGpt "peter-bird.com/chatgpt"
	)
	
	func main() {
		// Get the executable name
		appName := filepath.Base(os.Args[0])
		fmt.Println("Application Name:", appName)
	
		// Initialize Chat
		client, err := initChat()
		if err != nil {
			fmt.Printf("Error initializing API: %%s\n", err)
			return
		}
	
		// Setup the GUI
		myApp := app.New()
		myWin := myApp.NewWindow(appName + " App")
		setupGUI(myWin, client)
	
		myWin.ShowAndRun()
	}
	
	func initChat() (*chatGpt.ChatGPTClient, error) {
	
		pkgs := "/home/julian/go/pkgs"
	
		apikey := apiKey.NewAPIKey()
		key := apikey.ReadAPIKey(pkgs + "/apikey/config.json")
		fmt.Println(key)
	
		client, err := chatGpt.NewChatGPT(pkgs+"/chatgpt/config.json", key)
		if err != nil {
			return nil, fmt.Errorf("NewChatGPT returned an error: %` + `s", err)
		}
	
		if client == nil {
			return nil, fmt.Errorf("Expected a non-nil ChatGPTClient")
		}
	
		return client, nil
	}
	
`

	return temp
}

func (s *Source) GenGui() string {

	temp := `package main

	import (
		"fmt"
		"time"
	
		"fyne.io/fyne/v2"
		"fyne.io/fyne/v2/container"
		"fyne.io/fyne/v2/layout"
		"fyne.io/fyne/v2/widget"
	
		chatGpt "peter-bird.com/chatgpt"
	)
	
	func setupGUI(window fyne.Window, client *chatGpt.ChatGPTClient) {
		label := widget.NewLabel("Ask GPT")
	
		inputEntry, outputEntry, timerLabel, progress := setupInputOutput()
	
		sendButton := setupSendButton(inputEntry, outputEntry, timerLabel, progress, client)
		clsButton := widget.NewButton("Clear", func() {
			inputEntry.SetText("")
			outputEntry.SetText("")
		})
	
		myContainer := container.NewVBox(
			label,
			inputEntry,
			container.NewHBox(sendButton, layout.NewSpacer(), progress, timerLabel, layout.NewSpacer(), clsButton),
			outputEntry,
		)
	
		window.SetContent(myContainer)
		window.Resize(fyne.NewSize(700, 650))
	}
	
	func setupInputOutput() (*widget.Entry, *widget.Entry, *widget.Label, *widget.ProgressBarInfinite) {
		inputEntry := widget.NewMultiLineEntry()
		inputEntry.SetPlaceHolder("Type your message")
		inputEntry.SetMinRowsVisible(5)
		inputEntry.Wrapping = fyne.TextWrapWord
	
		outputEntry := widget.NewMultiLineEntry()
		outputEntry.SetPlaceHolder("ChatGPT responses will appear here")
		outputEntry.SetMinRowsVisible(24)
		outputEntry.Wrapping = fyne.TextWrapWord
	
		timerLabel := widget.NewLabel("00:00:00")
		progress := widget.NewProgressBarInfinite()
		progress.Stop()
	
		return inputEntry, outputEntry, timerLabel, progress
	}
	
	func setupSendButton(inputEntry, outputEntry *widget.Entry, timerLabel *widget.Label, progress *widget.ProgressBarInfinite, client *chatGpt.ChatGPTClient) *widget.Button {
		return widget.NewButton("Send", func() {
			progress.Start()
			defer progress.Stop()
	
			start := time.Now()
			ticker := time.NewTicker(time.Second)
	
			prompt := inputEntry.Text
			outputEntry.Append("User: " + prompt + "\n")
	
			fmt.Println(prompt)
	
			res, err := client.GetGPTResponse("You are a general consultant", outputEntry.Text)
			if err != nil {
				fmt.Printf("ChatGPT returned an error: %` + `s", err)
			}
			outputEntry.Append("\nChatGPT: " + res + "\n\n")
	
			ticker.Stop()
	
			duration := time.Since(start)
			timerLabel.SetText(fmt.Sprintf("%` + `02d:%` + `02d:%` + `02d", int(duration.Hours()), int(duration.Minutes())%60, int(duration.Seconds())%60))
			timerLabel.Refresh()
		})
	}
`
	return temp
}
