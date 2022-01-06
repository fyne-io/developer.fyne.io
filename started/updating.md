---
layout: page
title: Updating Content
---


## Updating content in your GUI
---

Having completed the [hello world](/started/hello) tutorial or other
examples you will have created a basic user interface. In this page
we see how the content of a GUI can be updated from your code.

The first step is to assign the widget you want to update to a
variable. In the hello world tutorial we passed `widget.NewLabel`
directly into `SetContent()`, to update it we change that to two
different lines, such as:

```go
	clock := widget.NewLabel("")
	w.SetContent(clock)
```

Once the content has been assigned to a variable we can call functions
like `SetText("new text")`. For our example we will set the
content of our label to the current time, with the help of
`Time.Format`.

```go
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
```

That is all we need to do to change content of a visible item (see below for the full code).
However, we can go further and update content on a regular basis.

## Running in the background

Most applications will need to have processes that run in the background,
for example downloading data or responding to events.
To simulate this we will extend the code above to run every second.

Like with most go code we can create a goroutine (using the `go`
keyword) and run our code there. If we move the text update code to
a new function it can be called on initial display as well as
on a timer for regular updating. By combining a goroutine and the
`time.Tick` inside a for loop we can update the label every second.

```go
	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()
```

It is important to place this code before `ShowAndRun` or `Run` calls
because they will not return until the application closes.
With all of this together the code will run and update the user interface
each second, creating a basic clock widget.
The full code is as follows:

```go
package main

import (
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
}

func main() {
	a := app.New()
	w := a.NewWindow("Clock")

	clock := widget.NewLabel("")
	updateTime(clock)

	w.SetContent(clock)
	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()
	w.ShowAndRun()
}
```

{% include youtube.html id="h2ZOdTA3ew4" %}
