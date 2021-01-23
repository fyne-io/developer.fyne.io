package main

import (
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	//"fyne.io/fyne/v2/theme"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Button Widget")

	content := widget.NewButton("click me", func() {
		log.Println("tapped")
	})

	//content := widget.NewButtonWithIcon("Home", theme.HomeIcon(), func() {
	//	log.Println("tapped home")
	//})

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
