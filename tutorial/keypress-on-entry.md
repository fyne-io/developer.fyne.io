---
layout: page
title: Handle Key Presses in Entry

redirect_from: /develop/keypress-on-entry.html
---

In the traditional sense, GUI programs have used callbacks to customize actions for widgets. Fyne does not expose inserting custom callbacks to capture events on widgets, but it does not need to. The Go language is plenty extensible to make this work.

Instead we can simply use Type Embedding and extend the widget, and add a method to run when the escape key is pressed.

First create a new type struct, we will call it `escapeEntry`.

```go
type escapeEntry struct {
    widget.Entry
}
```

Next up we need to create our `onEnter` method for the `enterEntry` struct. This will later be used to run when the enter key is pressed to print out the written text and then clear the text.

```go
func (e * escapeEntry) onEsc() {
    fmt.Println(e.Entry.Text)
    e.Entry.SetText("")
}
```

As mentioned in [Extending existing widgets](https://developer.fyne.io/tutorial/extending-widgets), we follow good practice and create a constructor function and extend the `BaseWidget`.

```go
func newEscapeEntry() * escapeEntry {
    entry := & escapeEntry{}
    entry.ExtendBaseWidget(entry)
    return entry
}
```

Then override the `TypedKey()` method that's part of the `fyne.Focusable` interface,
this will allow us to intercept the standard key handling and pass through if we want.
Inside this method, we will use a conditional to check if the `key.Name` value matches the `Escape` key using the `fyne.KeyEscape` variable and if that is the case, we run our `onEsc` method.
We need to have a default case that calls `e.Entry.TypedKey(key)` to maintain the behavior of other keys in the entry.
This implementation can easily be extended to check for other keys in the future if necessary.

```go
func (e *escapeEntry) TypedKey(key *fyne.KeyEvent) {
    switch key.Name {
    case fyne.KeyEscape:
        e.onEsc()
    default:
        e.Entry.TypedKey(key)
    }
}
```

With this, the entry now performs a custom action that will be executed when the `TypedKey()` event occurs in our GUI program.

In the end, the resulting code could look something like this:

```go
package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

type escapeEntry struct {
	widget.Entry
}

func (e *escapeEntry) onEsc() {
	fmt.Println(e.Entry.Text)
	e.Entry.SetText("")
}

func newEscEntry() *escapeEntry {
	entry := &escapeEntry{}
	entry.ExtendBaseWidget(entry)
	return entry
}

func (e *escapeEntry) TypedKey(key *fyne.KeyEvent) {
	switch key.Name {
	case fyne.KeyEscape:
		e.onEsc()
	default:
		e.Entry.TypedKey(key)
	}
}

func main() {
	a := app.New()
	w := a.NewWindow("Messenger")

	entry := newEscEntry()

	w.SetContent(entry)
	w.ShowAndRun()
}
```

