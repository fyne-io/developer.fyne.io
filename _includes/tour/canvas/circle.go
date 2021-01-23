package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Circle")

	circle := canvas.NewCircle(color.White)
	circle.StrokeColor = color.Gray{0x99}
	circle.StrokeWidth = 5
	w.SetContent(circle)

	w.Resize(fyne.NewSize(100, 100))
	w.ShowAndRun()
}
