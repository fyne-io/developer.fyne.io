package main

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	//"fyne.io/fyne/layout"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Container")

	text1 := canvas.NewText("Hello", color.White)
	text2 := canvas.NewText("There", color.White)
	text2.Move(fyne.NewPos(20, 20))
	text3 := canvas.NewText("World", color.White)
	container := fyne.NewContainer(text1, text2, text3)
	//	container := fyne.NewContainerWithLayout(layout.NewGridLayout(2),
	//		text1, text2, text3)

	myWindow.SetContent(container)
	myWindow.ShowAndRun()
}
