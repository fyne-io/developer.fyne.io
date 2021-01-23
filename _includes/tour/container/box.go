package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Entry Widget")

	content := container.NewVBox(
		widget.NewLabel("The top row of the VBox"),
		container.NewHBox(
			widget.NewLabel("Label 1"),
			widget.NewLabel("Label 2"),
		),
	)

	content.Add(widget.NewButton("Add more items", func() {
		content.Add(widget.NewLabel("Added"))
	}))

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
