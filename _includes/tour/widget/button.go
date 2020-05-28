package main

import (
	"log"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	//"fyne.io/fyne/theme"
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
