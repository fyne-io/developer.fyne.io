package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	//"fyne.io/fyne/v2/layout"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Container")

	text1 := canvas.NewText("Hello", color.White)
	text2 := canvas.NewText("There", color.White)
	text2.Move(fyne.NewPos(20, 20))
	text3 := canvas.NewText("World", color.White)
	content := container.NewWithoutLayout(text1, text2, text3)
	// content := container.New(layout.NewGridLayout(2), text1, text2, text3)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
