package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
)

var state State
var application fyne.App
var window fyne.Window

func init() {
	state = NewState()
	application = app.New()
	window = application.NewWindow("go-grep")
}

// TODO:
// (1) Menu for setttings (https://github.com/fyne-io/fyne/blob/master/cmd/fyne_demo/main.go)
//       - numberWorkers
//       - bufferSize

func main() {

	settingsMenuItem := fyne.NewMenuItem("Settings", func() {
		w := application.NewWindow("Settings")
		w.Resize(fyne.NewSize(480, 480))
		w.Show()
	})

	aboutMenuItem := fyne.NewMenuItem("About", func() {
		w := application.NewWindow("About")
		w.Resize(fyne.NewSize(480, 480))
		w.Show()
	})

	menu := fyne.NewMenu("Menu", settingsMenuItem, aboutMenuItem)
	mainMenu := fyne.NewMainMenu(menu)
	window.SetMainMenu(mainMenu)

	middle := createMiddleComponent()
	bottom := createBottomComponent()

	content := container.New(layout.NewBorderLayout(nil, bottom, nil, nil),
		bottom, middle)

	application.Settings().SetTheme(theme.DarkTheme())

	window.SetMaster()
	window.CenterOnScreen()
	window.SetContent(content)
	window.Resize(fyne.NewSize(900, 700))
	window.ShowAndRun()
}
