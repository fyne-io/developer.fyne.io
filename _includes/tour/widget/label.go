package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Label Widget")

	content := widget.NewLabel("text")

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
