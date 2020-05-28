package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"net/url"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")

	bugURL, _ := url.Parse("https://github.com/fyne-io/fyne/issues/new")
	myWindow.SetContent(widget.NewHyperlink("Report a bug", bugURL))

	myWindow.ShowAndRun()
}
