package main

import (
	"net/url"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")

	bugURL, _ := url.Parse("https://github.com/fyne-io/fyne/issues/new")
	myWindow.SetContent(widget.NewHyperlink("Report a bug", bugURL))

	myWindow.ShowAndRun()
}
