package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/theme"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Image")

	image := canvas.NewImageFromResource(theme.FyneLogo())
	// image := canvas.NewImageFromFile(fileName)
	// image := canvas.NewImageFromImage(src)
	image.FillMode = canvas.ImageFillOriginal
	w.SetContent(image)

	w.ShowAndRun()
}
