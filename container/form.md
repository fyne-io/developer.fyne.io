---
title: Form

redirect_from:
- /tour/layout/formlayout
---

The `layout.FormLayout` is like a 2 column [grid layout](/container/grid)
but tweaked to lay out forms in an application.
The height of each item will be the larger of the two minimum heights
in each row. The width of the left item will be the largest minimum
width of all items in the first column whilst the second item in each
row will expand to fill the space.

This layout is more typically used within the `widget.Form` (for validation, submit and cancel buttons, etc) but it can
also be used directly with `layout.NewFormLayout()` passed to the first
parameter of `container.New(...)`.

```go
package main

import (
	"image/color"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Form Layout")

	label1 := canvas.NewText("Label 1", color.Black)
	value1 := canvas.NewText("Value", color.White)
	label2 := canvas.NewText("Label 2", color.Black)
	value2 := canvas.NewText("Something", color.White)
	grid := container.New(layout.NewFormLayout(), label1, value1, label2, value2)
	myWindow.SetContent(grid)
	myWindow.ShowAndRun()
}
```
