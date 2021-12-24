---
title: Window Handling

redirect_from:
- /tour/basics/windows
---

Windows are created using `App.NewWindow()` and need to be shown using
the `Show()` function. The helper method `ShowAndRun()` on `fyne.Window`
allows you to show your window and run the application at the same time.

If you wish to show a second window you must only call the `Show()`
function. This is illustrated in the `showAnother` function.

```go
package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")
	myWindow.SetContent(widget.NewLabel("Hello"))

	go showAnother(myApp)
	myWindow.ShowAndRun()
}

func showAnother(a fyne.App) {
	time.Sleep(time.Second * 5)

	win := a.NewWindow("Shown later")
	win.SetContent(widget.NewLabel("5 seconds later"))
	win.Resize(fyne.NewSize(200, 200))
	win.Show()

	time.Sleep(time.Second * 2)
	win.Close()
}
```

By default a window will be the right size to show its content
by checking the `MinSize()` function (more on that in later examples).
You can set a larger size by calling the `Window.Resize()` function.

Be aware that the desktop environment may have constraints that cause
windows to be smaller than requested.
