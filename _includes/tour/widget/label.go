package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Label Widget")

	content := widget.NewLabel("text")

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
