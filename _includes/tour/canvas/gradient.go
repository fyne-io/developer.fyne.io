package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"image/color"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Gradient")

	gradient := canvas.NewHorizontalGradient(color.White, color.Transparent)
	//gradient := canvas.NewRadialGradient(color.White, color.Transparent)
	w.SetContent(gradient)

	w.Resize(fyne.NewSize(100, 100))
	w.ShowAndRun()
}
