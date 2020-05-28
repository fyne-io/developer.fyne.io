package main

import (
	"log"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Choice Widgets")

	check := widget.NewCheck("Optional", func(value bool) {
		log.Println("Check set to", value)
	})
	radio := widget.NewRadio([]string{"Option 1", "Option 2"}, func(value string) {
		log.Println("Radio set to", value)
	})
	combo := widget.NewSelect([]string{"Option 1", "Option 2"}, func(value string) {
		log.Println("Select set to", value)
	})

	myWindow.SetContent(widget.NewVBox(check, radio, combo))
	myWindow.ShowAndRun()
}
