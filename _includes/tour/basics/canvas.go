package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Canvas")
	myCanvas := myWindow.Canvas()

	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	text := canvas.NewText("Text", green)
	text.TextStyle.Bold = true
	myCanvas.SetContent(text)
	go changeContent(myCanvas)

	myWindow.Resize(fyne.NewSize(100, 100))
	myWindow.ShowAndRun()
}

func changeContent(c fyne.Canvas) {
	time.Sleep(time.Second * 2)

	blue := color.NRGBA{R: 0, G: 0, B: 180, A: 255}
	c.SetContent(canvas.NewRectangle(blue))

	time.Sleep(time.Second * 2)
	c.SetContent(canvas.NewLine(color.Gray{Y: 180}))

	time.Sleep(time.Second * 2)
	red := color.NRGBA{R: 0xff, G: 0x33, B: 0x33, A: 0xff}
	circle := canvas.NewCircle(color.White)
	circle.StrokeWidth = 4
	circle.StrokeColor = red
	c.SetContent(circle)

	time.Sleep(time.Second * 2)
	c.SetContent(canvas.NewImageFromResource(theme.FyneLogo()))
}
