package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Entry Widget")

	content := widget.NewVBox(widget.NewLabel("The top row of VBox"),
		widget.NewHBox(
			widget.NewLabel("Label 1"),
			widget.NewLabel("Label 2")))

	content.Append(widget.NewButton("Add more items", func() {
		content.Append(widget.NewLabel("Appended"))
	}))

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
