---
layout: page
title: Your First App

order: 20
---


## Create your first Fyne app
---

Having completed the steps in the [getting started](/started/) document you're ready to build your first app. To illustrate the process we will build a simple hello world application.

A simple app starts by creating an app instance with app.New() and then opening a window with app.NewWindow(). Then a widget tree is defined that is set as the main content with SetContent() on a window. The app UI is then shown by calling ShowAndRun() on the window.

<div id="hello__function">
<div style="text-align: left" markdown="1">
```go
package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
	))

	w.ShowAndRun()
}
```
</div>
</div>

The code above can be built using the command `go build hello.go` and then executed either by running the `hello` command or by double clicking the icon. You could also bypass the compiling step and just run the code directly using `go run hello.go`.

Either approach will show a window that looks just like this:

<img src="/images/architecture/hello-normal.png" width="207" alt="Hello Window" />

If you prefer a light theme then just set the environment variable `FYNE_THEME=light` and you'll get:

<img src="/images/architecture/hello-light.png" width="207" alt="Hello Light Theme" />

That's all there is to getting started. To learn more you can read the full
[API documentation](http://developer.fyne.io/api/).
