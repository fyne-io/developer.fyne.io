---
title: Border

redirect_from:
- /tour/layout/borderlayout
---

The border layout is possibly the most widely used to construct user
interfaces as it allows the positioning of items around a central element
which will expand to fill the space. To create a border layout you need
to pass the `fyne.CanvasObject`s that should be positioned in a border
position to the layout (as well as the container as usual). This
syntax is a little different to the other layouts but is basically just
`layout.NewBorderLayout(top, bottom, left, right)` as illustrated in
the example to the right.

Any items passed to the container that do not appear in specific border
locations will be positioned to the central area and will expand to
fill the space available. You can also pass `nil` to border parameters
that you wish to leave empty.

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
	myWindow := myApp.NewWindow("Border Layout")

	top := canvas.NewText("top bar", color.White)
	left := canvas.NewText("left", color.White)
	middle := canvas.NewText("content", color.White)
	content := container.New(layout.NewBorderLayout(top, nil, left, nil),
		top, left, middle)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
```

Note that all items in the center will expand to fill the space (as if
they were in a [`layout.MaxLayout`](maxlayout.html) container). To manage
the area yourself you can create a new `fyne.Container` (using `container.New()`) and use any
layout you wish.
