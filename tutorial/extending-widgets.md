---
layout: page
title: Extending Widgets

order: 10
redirect_from: /develop/extending-widgets.html
---

## Extending Existing Widgets
---

The standard Fyne widgets provide the minimum functionality and customisation
to support most use-cases. It may be required at certain times to have more
advanced functionality. Rather than have developers build their own widgets
it is possible to extend the existing ones.

For example we will extend the icon widget to support being tapped. To do this
we declare a new struct that embeds the `widget.Icon` type. We create a
constructor function that calls the important `ExtendBaseWidget` function.

```go
import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

type tappableIcon struct {
	widget.Icon
}

func newTappableIcon(res fyne.Resource) *tappableIcon {
	icon := &tappableIcon{}
	icon.ExtendBaseWidget(icon)
	icon.SetResource(res)

	return icon
}
```

We then add new functions to implement the `fyne.Tappable` interface, with
those functions added the new `Tapped` function will be called every time the
user taps our new icon.  The interface required has two functions,
`Tapped(*PointEvent)` and `TappedSecondary(*PointEvent)`, so we will add both.

```go
import "log"

func (t *tappableIcon) Tapped(_ *fyne.PointEvent) {
	log.Println("I have been tapped")
}

func (t *tappableIcon) TappedSecondary(_ *fyne.PointEvent) {
}
```

We can test this new widget using a simple application as follows.

```go
import (
    "fyne.io/fyne/app"
    "fyne.io/fyne/theme"
)

func main() {
	a := app.New()
	w := a.NewWindow("Tappable")
	w.SetContent(newTappableIcon(theme.FyneLogo()))
	w.ShowAndRun()
}
```

## Extending a Widget Renderer

We can also extend a widget renderer, allowing you to override its default 
appearance. For instance if we want to override the default background colour 
of the Entry widget, we first have to extend the Entry widget as follows.

```go
// create an extended Entry widget
type skinnedEntry struct {
	widget.Entry
}

func newSkinnedEntry() *skinnedEntry {
	skinnedEntry := &skinnedEntry{}
	skinnedEntry.ExtendBaseWidget(skinnedEntry)
	return skinnedEntry
}
```

We then need to extend the renderer as well. For our purposes we cache a
copy of the Entry widget renderer and pass through to that for the functions 
where we don't need to amend the behaviour.

```go
// Create the skinned entry renderer by caching the entry renderer and including
// a reference to ourselves
func (skinnedEntry *skinnedEntry) CreateRenderer() fyne.WidgetRenderer {
	skinnedEntry.ExtendBaseWidget(skinnedentry)
	entryRenderer := skinnedEntry.Entry.CreateRenderer()
	return &skinnedEntryRenderer{
		skinnedEntry: skinnedEntry,
		entryRenderer: entryRenderer,
	}
}

type skinnedEntryRenderer struct {
	skinnedEntry *skinnedEntry
	entryRenderer fyne.WidgetRenderer
}

// Pass through these functions to the base Entry widget renderer
func (renderer *skinnedEntryRenderer) Refresh() {
	renderer.entryRenderer.Refresh()
}

func (renderer *skinnedEntryRenderer) MinSize() fyne.Size {
	return renderer.entryRenderer.MinSize()
}

func (renderer *skinnedEntryRenderer) Layout(size fyne.Size) {
	renderer.entryRenderer.Layout(size)
}

func (renderer *skinnedEntryRenderer) Objects() []fyne.CanvasObject {
	return renderer.entryRenderer.Objects()
}

func (renderer *skinnedEntryRenderer) Destroy() {
  renderer.entryRenderer.Destroy()
}

// Override the background colour function to provide one explicitly, rather 
// than from the theme
func (renderer *skinnedEntryRenderer) BackgroundColor() color.Color {
	return color.NRGBA{R: 0x7c, G: 0x7c, B: 0x7c, A: 0xff}
}
```

We can test out our new extended widget using a simple program such as the below.

```go
import (
	"image/color"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Skinned Entry")
	w.SetContent(newSkinnedEntry())
	w.ShowAndRun()
}
``` 

The result is that a window will appear with a gray-coloured entry box instead 
of whatever colour is the usual for the theme. 
