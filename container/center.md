---
title: Center

redirect_from:
- /tour/layout/centerlayout
---

`layout.CenterLayout` organises all items in its container to be
centered in the available space. The objects will be drawn in the order
they are passed to the container, with the last being drawn top-most.

```go
package main

import (
	"image/color"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Center Layout")

	img := canvas.NewImageFromResource(theme.FyneLogo())
	img.FillMode = canvas.ImageFillOriginal
	text := canvas.NewText("Overlay", color.Black)
	content := container.New(layout.NewCenterLayout(), img, text)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
```

The center layout causes all items to stay at their minimum size, if
you wish to expand items to fill the space then see
[`layout.MaxLayout`](max).


Example: 

<img width="970" alt="image" src="https://user-images.githubusercontent.com/96031819/204343893-0dcfc9f7-de71-47e2-9b26-489b92da99c3.png">
